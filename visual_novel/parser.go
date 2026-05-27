package visual_novel

import "fmt"

func CheckValid(script []ScriptEntry) {
	for _, entry := range script {
		fmt.Printf("ID %d: Type=%s Action=%s\n", entry.ID, entry.Type, entry.Action)

		switch entry.Type {
		case "dialogue":
			fmt.Printf("%s says: %s\n", entry.Label, entry.Content)
		case "show_sprite":
			fmt.Printf("Show sprite %s at %s\n", entry.Sprite, entry.Location)
		case "remove_sprite":
			fmt.Printf("Remove sprite %s\n", entry.Sprite)
		case "game_action":
			fmt.Printf("Game action: %s\n", entry.Actions)
		}
	}

	fmt.Println("extracted")
}
