# EDITORJS GO CODE

`editorjs-go-code` is a plugin for [`editorjs-go`](https://github.com/christiandenisi/editorjs-go) that renders Editor.js `code` blocks using `<pre><code>` with safe HTML escaping.

## ✨ Features

- ✅ Escapes code content for safe rendering
- ✅ Semantic HTML output: `<pre><code>...</code></pre>`
- 🧩 Easy integration with `editorjs-go`

---

## 📦 Installation

```bash
go get github.com/christiandenisi/editorjs-go
go get github.com/christiandenisi/editorjs-go-code
```

---

## 🚀 Usage

```go
package main

import (
    "fmt"
    "github.com/christiandenisi/editorjs-go"
    "github.com/christiandenisi/editorjs-go-code"
)

func main() {
    jsonData := []byte(`{
        "blocks": [{
            "type": "code",
            "data": {
                "code": "<script>alert('x')</script>"
            }
        }]
    }`)

    conv := editorjs.New()
    editorjs.Register(conv, "code", code.RenderCode)

    html, err := conv.Convert(jsonData)
    if err != nil {
        panic(err)
    }

    fmt.Println(html)
}
```

---

## 📌 Notes

- All code content is escaped using `html.EscapeString` to prevent HTML injection.
- Suitable for displaying code in documentation or blogs.

---

## 🧱 Compatibility

- Compatible with `editorjs-go` version 1.x

---

## 👤 License

MIT