package code

import (
    "testing"
    "github.com/christiandenisi/editorjs-go"
)

func TestRenderCode(t *testing.T) {
    block := editorjs.Block[CodeData]{
        Data: CodeData{
            Code: "<script>alert('x')</script>",
        },
    }

    ctx := &editorjs.Context{}
    out, err := RenderCode(block, ctx)
    if err != nil {
        t.Fatalf("RenderCode returned an error: %v", err)
    }

    expected := "<pre><code>&lt;script&gt;alert(&#39;x&#39;)&lt;/script&gt;</code></pre>"
    if out != expected {
        t.Errorf("unexpected output: got %q, want %q", out, expected)
    }
}