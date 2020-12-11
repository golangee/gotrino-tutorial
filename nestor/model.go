package nestor

import (
	"strings"
	"unicode"
)

// A Fragment is a semantically parsed "tutorial.md" file.
type Fragment struct {
	Title       string
	Teaser      []*Attachment `json:",omitempty"`
	Body        string
	Attachments []*Attachment `json:",omitempty"`
	Fragments   []*Fragment   `json:",omitempty"` // Fragments contain more children fragments
}

// ID creates a string id from the title.
func (f *Fragment) ID() string {
	return strings.ToLower(text2GoIdentifier(f.Title))
}

// AttachmentType denotes the semantically parsed element.
type AttachmentType string

const (
	AtImage    AttachmentType = "image"
	AtVideo    AttachmentType = "video"
	AtSource   AttachmentType = "source"
	AtIFrame   AttachmentType = "iframe"
	AtDownload AttachmentType = "download"
)

// Attachment represents an entry of the Attachment section.
type Attachment struct {
	Type AttachmentType
	// Title may be only non-empty for image or video
	Title string

	// Raw is usually a relative or absolute URL. The semantic is duck-typed, so if it can be found nearby
	// locally, File contains the absolute Filepath.
	Raw string

	// File contains the absolute local filepath, if available.
	File string
}

func text2GoIdentifier(p string) string {
	sb := &strings.Builder{}
	upCase := true
	written := 0
	for _, r := range p {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			upCase = true
			continue
		}

		if r >= '0' && r <= '9' && written == 0 {
			sb.WriteRune('S')
		}

		written++
		if upCase {
			sb.WriteRune(unicode.ToUpper(r))
			upCase = false
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
