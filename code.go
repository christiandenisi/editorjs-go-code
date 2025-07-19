package code

import (
    "html"

    "github.com/christiandenisi/editorjs-go"
)

// CodeData represents the data structure for a code block in Editor.js.
type CodeData struct {
    Code string `json:"code"`
}

// RenderCode is the renderer function for the "code" block.
// It escapes code content and wraps it inside <pre><code>...</code></pre>.
func RenderCode(b editorjs.Block[CodeData], ctx *editorjs.Context) (string, error) {
    escaped := html.EscapeString(b.Data.Code)
    return "<pre><code>" + escaped + "</code></pre>", nil
}