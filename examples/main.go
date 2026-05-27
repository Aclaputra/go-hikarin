package main

import "github.com/aclaputra/go-hikarin/visual_novel"

func main() {
	script, err := visual_novel.ExtractFromFile("./scripts/example_script.json")
	if err != nil {
		panic(err)
	}

	visual_novel.CheckValid(script)
}
