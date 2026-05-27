# Visual Novel Script Library (Go)

A Go library for parsing and executing **Hikarin Framework** visual novel scripts.  
Inspired by the [`VisualNovelEngine.java`](https://github.com/Iteranya/MinecraftHikarinMod/blob/master/src/main/java/org/arsparadox/mobtalkerredux/vn/controller/VisualNovelEngine.java) from the open‑source [MinecraftHikarinMod](https://github.com/Iteranya/MinecraftHikarinMod), this project translates the same ideas into idiomatic Go.

The JSON script format is exported from the [Iteranya/hikarin-framework](https://github.com/Iteranya/hikarin-framework), a platform‑agnostic scripting system for visual novels.

---

## ✨ Features
- Parse Hikarin JSON scripts into Go structs (`ScriptEntry`)
- Iterate through script entries one by one
- Handle dialogue, sprite management, labels, and game actions
- Flexible API: load from file path, raw bytes, or any `io.Reader`
- Designed as a reusable library for embedding into Go game engines (e.g. [Ebitengine](https://ebiten.org))

---

## 📦 Installation
```bash
go get github.com/aclaputra/visualnovel/vnlib
```

## 📂 Folder Structure
```
visualnovel/
├── go.mod
├── vnlib/
│   ├── model.go        # ScriptEntry struct
│   ├── extract.go      # ExtractFromBytes, ExtractFromFile, ExtractFromReader
│   └── parser.go       # Helpers for handling entries
├── scripts/
│   └── status_quo_script.json   # Example script
└── examples/
    └── main.go         # Demo program
```

## 🧑‍💻 Usage
## Example Script
```json
[
  { "type": "dialogue", "action": "say", "label": "Conductor", "content": "Evening again, miss.", "id": 3 },
  { "type": "remove_sprite", "action": "remove_character", "sprite": "conductor", "id": 4 }
]
```
`See more on examples and scripts folder`

## 📖 Inspirations
MinecraftHikarinMod (github.com) – original Java implementation of a VN engine.

Hikarin Framework – JSON scripting format exported and consumed by this library.

## ⚖️ License
This library is open source under the MIT License.
It builds upon ideas from the Hikarin Framework and MinecraftHikarinMod, both of which are open source projects.

## 🚀 Roadmap
[ ] Add support for branching choices

[ ] Integrate with Ebitengine rendering

[ ] Provide unit tests for script parsing

[ ] Expand documentation with advanced examples