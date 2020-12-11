package nestor

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const magicFilename = "tutorial.md"

// ParseDocument takes the buffer and interprets it according to our semantic Markdown specification.
func ParseDocument(buf []byte) Fragment {
	lines := strings.Split(string(buf), "\n")

	const (
		prefixAttachments = "## Attachments"
	)

	var frag Fragment
	const (
		stateFindTitle = iota
		stateFindBody
		stateFindAttachments
	)
	state := stateFindTitle
	tmpBody := &strings.Builder{}

	for _, line := range lines {
		tLine := strings.TrimSpace(line)
		switch state {
		case stateFindTitle:
			if strings.HasPrefix(tLine, "# ") {
				frag.Title = strings.TrimSpace(tLine[1:])
				state = stateFindBody
			}
		case stateFindBody:
			if name, value := parseLink(tLine); name == "teaser" && strings.HasPrefix(tLine, "!") {
				frag.Teaser = append(frag.Teaser, &Attachment{
					Type:  atTypeFromName(value),
					Title: name,
					Raw:   value,
				})
			} else {
				if strings.HasPrefix(tLine, prefixAttachments) {
					state = stateFindAttachments
					continue
				}

				tmpBody.WriteString(line)
				tmpBody.WriteString("\n")
			}
		case stateFindAttachments:
			if name, value := parseLink(tLine); name != "" && strings.HasPrefix(tLine, "!") {
				frag.Attachments = append(frag.Attachments, &Attachment{
					Type:  atTypeFromName(value),
					Title: name,
					Raw:   value,
				})
			} else

			if name, value := parseLink(tLine); name == "iframe" {
				frag.Attachments = append(frag.Attachments, &Attachment{Type: AtIFrame, Title: name, Raw: value})
			} else
			if name, value := parseLink(tLine); name == "download" {
				frag.Attachments = append(frag.Attachments, &Attachment{Type: AtDownload, Title: name, Raw: value})
			} else
			if name, value := parseLink(tLine); name == "source" {
				frag.Attachments = append(frag.Attachments, &Attachment{Type: AtSource, Title: name, Raw: value})
			}

		}
	}

	frag.Body = strings.TrimSpace(tmpBody.String())
	return frag
}

// ParseFile reads the given file name and tries to resolve the attachments to absolute filenames.
func ParseFile(file string) (*Fragment, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, nil
	}

	frag := ParseDocument(buf)

	resolveFiles(&frag, filepath.Dir(file))

	return &frag, nil
}

// ParseDir searches recursively and returns all found Fragment roots.
func ParseDir(root string) ([]*Fragment, error) {
	tmp := &Fragment{}
	if err := findChildren(tmp, root); err != nil {
		return nil, err
	}

	return tmp.Fragments, nil
}

func findChildren(parent *Fragment, dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// 1. find new parent
	for _, file := range files {
		if file.Mode().IsRegular() && file.Name() == magicFilename {
			p, err := ParseFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return fmt.Errorf("unable to parse file: %w", err)
			}

			parent.Fragments = append(parent.Fragments, p)
			parent = p
		}
	}

	// 2. find children
	for _, file := range files {
		if file.Mode().IsDir() && !strings.HasPrefix(file.Name(), ".") {
			if err := findChildren(parent, filepath.Join(dir, file.Name())); err != nil {
				return err
			}
		}
	}

	return nil
}

// atTypeFromName only distinguishes between images and videos. If it is not an image, it is a video.
func atTypeFromName(name string) AttachmentType {
	name = strings.ToLower(name)
	imgs := []string{".png", ".jpg", ".jpeg", ".gif", ".webp"}
	for _, img := range imgs {
		if strings.HasSuffix(name, img) {
			return AtImage
		}
	}

	return AtVideo
}

// parseLink takes any text. However, if the form is [*](*) it returns the according non-empty texts.
func parseLink(text string) (name, target string) {

	nStart := strings.IndexByte(text, '[')
	nEnd := strings.IndexByte(text, ']')

	tStart := strings.IndexByte(text, '(')
	tEnd := strings.IndexByte(text, ')')

	if nStart == -1 || nEnd == -1 || tStart == -1 || tEnd == -1 || nStart > nEnd || tStart > tEnd || nStart > tStart || nEnd > tEnd {
		return
	}

	return text[nStart+1 : nEnd], text[tStart+1 : tEnd]
}

func resolveFiles(frag *Fragment, dir string) {
	resolveFilesList(frag.Teaser, dir)
	resolveFilesList(frag.Attachments, dir)
}

func resolveFilesList(list []*Attachment, dir string) {
	for i := range list {
		// for security reason: we do not support ".." in local file links
		raw := list[i].Raw
		if strings.Contains(raw, "..") {
			list[i].File = raw
		} else {
			fname := filepath.Join(dir, raw)
			if _, err := os.Stat(fname); err != nil {
				list[i].File = raw
			} else {
				list[i].File = fname
			}
		}
	}
}
