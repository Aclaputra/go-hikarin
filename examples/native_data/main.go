package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aclaputra/go-hikarin/hikarin"
)

func main() {
	rawItems, err := hikarin.ExtractRaw("./scripts/example_script_conditional.json")
	if err != nil {
		log.Fatal(err)
	}

	Parse(rawItems, 0, false)
}

var GlobalValues = make(map[string]int, 0)

// replace fmt println with actions needed in your game or app
func Process(rawItems []json.RawMessage, currentState uint, stopRecursive bool) {
	for _, rawItem := range rawItems {
		var finder struct {
			Type string `json:"type"`
			Id   int    `json:"id"`
		}
		if err := json.Unmarshal(rawItem, &finder); err != nil {
			log.Println("Error parsing type:", err)
			continue
		}
		skipIfNotCurrentId := finder.Id != int(currentState)
		if skipIfNotCurrentId {
			continue
		}

		switch finder.Type {
		case hikarin.META:
			var m hikarin.Meta
			if err := json.Unmarshal(rawItem, &m); err == nil {
				fmt.Println("> Meta", m.Action)
				if m.Action == "start" {
					Parse(rawItems, currentState, true)
				}
			}
		case hikarin.LABEL:
			var l hikarin.Label
			if err := json.Unmarshal(rawItem, &l); err == nil {
				fmt.Println("> Label", l.Label)
			}
		case hikarin.SHOW_SPRITE:
			var ss hikarin.ShowSprite
			if err := json.Unmarshal(rawItem, &ss); err == nil {
				fmt.Println("> Show Sprite")
				fmt.Println("Set sprite", ss.Sprite)
				fmt.Println("Set Location", ss.Location)

				if ss.Action == "show" {
					fmt.Println("wRatio", ss.WRatio)
					fmt.Println("hRatio", ss.HRatio)
					fmt.Println("wFrameRatio", ss.WFrameRatio)
					fmt.Println("hFrameRatio", ss.HFrameRatio)
					fmt.Println("column", ss.Column)
					fmt.Println("row", ss.Row)
				}
			}
		case hikarin.DIALOGUE:
			var d hikarin.Dialogue
			if err := json.Unmarshal(rawItem, &d); err == nil {
				fmt.Println("> Dialogue")
				fmt.Println("Set Label", d.Label)
				fmt.Println("Set Content", d.Content)
			}
		case hikarin.MODIFY_GLOBAL:
			var md hikarin.ModifyGlobal
			if err := json.Unmarshal(rawItem, &md); err == nil {
				fmt.Println("> Modify Global")
				fmt.Println("Get Var name", md.Var)
				fmt.Println("Get Action", md.Action)
				fmt.Println("Get Value", md.Value)
				// GlobalValues[md.Var] = md.Value
				if md.Action == "increment_var" {
					GlobalValues[md.Var]++
					fmt.Printf(`%s value of %d increased to %d`, md.Var, md.Value, GlobalValues[md.Var])
					fmt.Println()
				}
			}
		case hikarin.REMOVE_SPRITE:
			var rs hikarin.RemoveSprite
			if err := json.Unmarshal(rawItem, &rs); err == nil {
				fmt.Println("> Remove Sprite")
				fmt.Println("Action", rs.Action)
				fmt.Println("Remove Sprite", rs.Sprite)
			}
		case hikarin.CHOICE:
			var c hikarin.Choice
			if err := json.Unmarshal(rawItem, &c); err == nil {
				fmt.Println("> Choice")
				for _, choice := range c.Choice {
					fmt.Println(choice.Display, "then go to label", choice.Label)
				}
			}
		case hikarin.CONDITIONAL_GLOBAL:
			var cg hikarin.ConditionalGlobal
			if err := json.Unmarshal(rawItem, &cg); err == nil {
				fmt.Println("> Conditional Global")
				fmt.Println("Get ondition", cg.Condition)
				fmt.Println("Get Var name", cg.Var)
				fmt.Println("Get Value", cg.Value)
				fmt.Println("Get End", cg.End)
				global := int64(GlobalValues[cg.Var])
				switch cg.Condition {
				case "equal":
					if global == cg.Value {
						fmt.Println("> Go to Actions - conditions met")
						Parse(cg.Actions, currentState, true)
					}
				case "not_equal":
					if global != cg.Value {
						fmt.Println("> Go to Actions - conditions met")
						Parse(cg.Actions, currentState, true)
					}
				case "less_than":
					if global < cg.Value {
						fmt.Println("> Go to Actions - conditions met")
						Parse(cg.Actions, currentState, true)
					}
				case "greater_than":
					if global > cg.Value {
						fmt.Println("> Go to Actions - conditions met")
						Parse(cg.Actions, currentState, true)
					}
				}
			}
		case hikarin.GAME_ACTION:
			var ga hikarin.GameAction
			if err := json.Unmarshal(rawItem, &ga); err == nil {
				fmt.Println("> Game Action")
				fmt.Println("get actions", ga.Actions)
			}
		default:
			log.Printf("Unknown type: %s\n", finder.Type)
		}
		log.Println("Next")
		fmt.Println("CurrentState:", currentState)
		fmt.Println()
		if stopRecursive {
			return
		}
		currentState++ // increment
	}
}

func Parse(rawItems []json.RawMessage, currentState uint, stopRecursive bool) {
	for _, rawItem := range rawItems {
		var finder struct {
			Type string `json:"type"`
			Id   int    `json:"id"`
		}

		if err := json.Unmarshal(rawItem, &finder); err != nil {
			log.Println("Error parsing type:", err)
			continue
		}
		skipIfNotCurrentId := finder.Id != int(currentState)
		if skipIfNotCurrentId {
			continue
		}

		switch finder.Type {
		case "meta":
			var m hikarin.Meta
			if err := json.Unmarshal(rawItem, &m); err == nil {
				switch m.Action {
				case "create_global": // global
					GlobalValues[m.Var] = int(m.Init)
					fmt.Println("> Create Global", m.Var, m.Init)
					currentState++
				case "create_var": // local
					// better separate the local and global
					GlobalValues[m.Var] = int(m.Init)
					fmt.Println("> Create Local", m.Var, m.Init)
					currentState++
				case "start":
					currentState++
					fmt.Println("> Start Process")
					Process(rawItems, currentState, stopRecursive)
				}
			}
		}
	}
}
