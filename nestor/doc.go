// Package nestor provides a markdown tutorial engine, which provides a simple semantic data model.
// The core declaration is based on standard Markdown, organized in so-called fragments, which is
// just a "tutorial.md", with some extra semantics. Example fragment pattern:
//
//   # Title of something             <- the fragments title, "# *" is the pattern
//   ![teaser](teaser.jpg)            <- the fragments teaser image, "![teaser](*)" is the pattern
//   Hello and Goodbye.               <- normal body text
//                                    <- normal body break
//   Just a usual *markdown* text.    <- normal body text
//
//   ## Attachments                   <- additional information, "## Attachments" is the pattern
//
//   ![Download Go](how-to-click.png) <- image or video, "![*](*) is the pattern
//
//   [iframe](00-setup)               <- iframe for a live preview, "[iframe](*)" is the pattern
//
//   [source](component.go)           <- embedding a source view, "[source](*)" is the pattern
//
//   [download](project.zip)          <- referring a download, "[download](*)" is the pattern
package nestor
