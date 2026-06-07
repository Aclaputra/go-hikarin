package hikarin

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Extract using raw original data structure of hikarin framework
func ExtractRaw(path string) (rawItems []json.RawMessage, err error) {
	var bytesData []byte
	bytesData, err = os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = json.Unmarshal(bytesData, &rawItems); err != nil {
		log.Fatal(err)
		return
	}

	return
}

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
