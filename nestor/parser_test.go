package nestor

import (
	"reflect"
	"testing"
)

func Test_parseLink(t *testing.T) {
	tests := []struct {
		name       string
		args       string
		wantName   string
		wantTarget string
	}{
		{"1", "", "", ""},
		{"2", "[]()", "", ""},
		{"3", "[a](b)", "a", "b"},
		{"4", "[ab](bc)", "ab", "bc"},
		{"4", " abc  [ab](cd) xyz ", "ab", "cd"},
		{"4", "[ab(bc)", "", ""},
		{"3", "(a)[b]", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotTarget := parseLink(tt.args)
			if gotName != tt.wantName {
				t.Errorf("parseLink() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotTarget != tt.wantTarget {
				t.Errorf("parseLink() gotTarget = %v, want %v", gotTarget, tt.wantTarget)
			}
		})
	}
}

func TestParseDocument(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want Fragment
	}{
		{
			name: "empty",
			args: []byte(""),
			want: Fragment{},
		},

		{
			name: "title only",
			args: []byte("# fragment Title"),
			want: Fragment{
				Title: "fragment Title",
			},
		},

		{
			name: "title + body",
			args: []byte("# fragment Title\nHey body."),
			want: Fragment{
				Title: "fragment Title",
				Body:  "Hey body.",
			},
		},

		{
			name: "title + body + lines",
			args: []byte("\n # fragment Title\n\nHey body.\n\n"),
			want: Fragment{
				Title: "fragment Title",
				Body:  "Hey body.",
			},
		},

		{
			name: "title+teaser+body",
			args: []byte("# fragment Title\n\n![teaser](image.jpg)\n\nHey body."),
			want: Fragment{
				Title: "fragment Title",
				Body:  "Hey body.",
				Teaser: []Attachment{
					{
						Type:  AtImage,
						Title: "teaser",
						Raw:   "image.jpg",
					},
				},
			},
		},

		{
			name: "title+teaser+body+attachments",
			args: []byte("# fragment Title\n\n![teaser](image.jpg)\n\nHey body.\n\n## Attachments\n\n" +
				"![Download Go](how-to-click.png)\n\n" +
				"[iframe](00-setup)  \n\n" +
				"  [source](component.go)\n\n" +
				"[download](project.zip)"),
			want: Fragment{
				Title: "fragment Title",
				Body:  "Hey body.",
				Teaser: []Attachment{
					{
						Type:  AtImage,
						Title: "teaser",
						Raw:   "image.jpg",
					},
				},
				Attachments: []Attachment{
					{
						Type:  AtImage,
						Title: "Download Go",
						Raw:   "how-to-click.png",
					},
					{
						Type:  AtIFrame,
						Title: "iframe",
						Raw:   "00-setup",
					},
					{
						Type:  AtSource,
						Title: "source",
						Raw:   "component.go",
					},
					{
						Type:  AtDownload,
						Title: "download",
						Raw:   "project.zip",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDocument(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDocument() = \n%+v, want \n%+v", got, tt.want)
			}
		})
	}
}
