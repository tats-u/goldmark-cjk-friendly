# goldmark-cjk-friendly

This is a port of [`remark-cjk-friendly` / `markdown-it-cjk-friendly`](https://github.com/tats-u/markdown-cjk-friendly) for [Goldmark](https://github.com/yuin/goldmark) as an external extension.

## Usage

```go
package main

import (
    "github.com/yuin/goldmark"
    "github.com/tats-u/goldmark-cjk-friendly"
)

func main() {
    md := goldmark.New(
        goldmark.WithExtensions(
            cjkfriendly.CJKFriendlyEmphasis,
        ),
    )
}
```

- `CJKFriendlyEmphasis`: The basic extension without GFM strikethrough support
- `CJKFriendlyEmphasisAndStrikethrough`: The basic extension with GFM strikethrough support. You do not need to add Goldmark's `Strikethrough` extension if you use this extension.
- `CJKFriendlyStrikethrough`: `CJKFriendlyEmphasisAndStrikethrough` - `CJKFriendlyEmphasis`. However, you do not need to prefer this extension to `CJKFriendlyEmphasisAndStrikethrough` since you will want to use `CJKFriendlyEmphasis` and `CJKFriendlyStrikethrough` together. Use this extension only if you want to switch between this extension and Goldmark's plain `Strikethrough` extension.

## Comparison with Goldmark's `EscapedSpace` extension

Combining this extension with Goldmark's `EscapedSpace` extension is welcome. They are not mutually exclusive. If you meet a case that cannot be emphasized with this extension, you can rely on Goldmark's `EscapedSpace` extension:

```md
a\ **()**\ aあ**()**あ
```

This extension can handle the case `あ**()**あ`, but cannot handle the case `a**()**a` for the compatibility with the plain CommonMark. You can add `\ ` outside a emphasis expression to make `EscapedSpace` extension work.

```go
package main

import (
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/extension"
    "github.com/tats-u/goldmark-cjk-friendly"
)

func main() {
    md := goldmark.New(
        goldmark.WithExtensions(
            cjkfriendly.CJKFriendlyEmphasis,
            extension.NewCJK(extension.WithEscapedSpace()), // or extension.CJK,
        ),
    )
}
```

## License

MIT (same as Goldmark)