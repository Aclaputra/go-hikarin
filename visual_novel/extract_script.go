package visual_novel

import (
	"encoding/json"
	"io"
	"os"
)

func ExtractFromFile(path string) ([]ScriptEntry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ExtractFromBytes(data)
}

func ExtractFromBytes(data []byte) ([]ScriptEntry, error) {
	var script []ScriptEntry
	if err := json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

func ExtractFromReader(r io.Reader) ([]ScriptEntry, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ExtractFromBytes(data)
}
