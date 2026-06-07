package main

import "github.com/aclaputra/go-hikarin/hikarin"

func main() {
	script, err := hikarin.ExtractFromFile("./scripts/example_script.json")
	if err != nil {
		panic(err)
	}

	hikarin.CheckValid(script)
}
