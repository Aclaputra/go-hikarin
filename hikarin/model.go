package hikarin

import (
	"encoding/json"
	"fmt"
	"log"
)

type ScriptEntry struct {
	Type        string      `json:"type"`
	Action      string      `json:"action"`
	ID          int         `json:"id"`
	Label       string      `json:"label,omitempty"`
	Sprite      string      `json:"sprite,omitempty"`
	Location    string      `json:"location,omitempty"`
	DynLocation string      `json:"dyn_location,omitempty"`
	Position    string      `json:"position,omitempty"`
	WRatio      int         `json:"wRatio,omitempty"`
	HRatio      int         `json:"hRatio,omitempty"`
	WFrameRatio int         `json:"wFrameRatio,omitempty"`
	HFrameRatio int         `json:"hFrameRatio,omitempty"`
	Column      int         `json:"column,omitempty"`
	Row         int         `json:"row,omitempty"`
	Content     string      `json:"content,omitempty"`
	Actions     interface{} `json:"actions,omitempty"`
}

type Meta struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	Var    string `json:"var,omitempty"`
	Init   int64  `json:"init,omitempty"`
	Id     int16  `json:"id"`
}

type Label struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	Label  string `json:"label,omitempty"`
	Id     int16  `json:"id"`
}

type ShowSprite struct {
	Type        string `json:"type"`
	Action      string `json:"action"`
	Sprite      string `json:"sprite"`
	Location    string `json:"location"`
	DynLocation string `json:"dyn_location"`
	Position    string `json:"position"`
	WRatio      int    `json:"wRatio"`
	HRatio      int    `json:"hRatio"`
	WFrameRatio int    `json:"wFrameRatio"`
	HFrameRatio int    `json:"hFrameRatio"`
	Column      int    `json:"column"`
	Row         int    `json:"row"`
	Id          int    `json:"id"`
}

type Dialogue struct {
	Type    string `json:"type"`
	Action  string `json:"action"`
	Label   string `json:"label"`
	Content string `json:"content"`
	Id      int    `json:"id"`
}

type ModifyGlobal struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	Var    string `json:"var"`
	Value  int    `json:"value"`
	Id     int    `json:"id"`
}

type RemoveSprite struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	Sprite string `json:"sprite"`
	Id     int    `json:"id"`
}

type Choice struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	Choice []ChoiceDetail
	Id     int `json:"id"`
}
type ChoiceDetail struct {
	Label   string `json:"label"`
	Display string `json:"display"`
}

type ConditionalGlobal struct {
	Type      string            `json:"type"`
	Action    string            `json:"action"`
	Condition string            `json:"condition"`
	Var       string            `json:"var"`
	Value     int64             `json:"value"`
	Actions   []json.RawMessage `json:"actions"`
	Id        int               `json:"id"`
	End       int               `json:"end"`
}
type GameAction struct {
	Type    string `json:"type"`
	Actions string `json:"actions"`
	Id      int    `json:"id"`
}

const (
	META               = "meta"
	LABEL              = "label"
	SHOW_SPRITE        = "show_sprite"
	DIALOGUE           = "dialogue"
	MODIFY_GLOBAL      = "modify_global"
	REMOVE_SPRITE      = "remove_sprite"
	CHOICE             = "choice"
	CONDITIONAL_GLOBAL = "conditional_global"
	GAME_ACTION        = "game_action"
)

/*
* WIP
LinkedList Data Structure
*/
type Node struct {
	Value int // the value can be struct
	Next  *Node
	Prev  *Node
}

// Doubly LinkedList for Dialogue Chains Data Structure
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// Append adds a new node with the given value to the end of the list
func (dll *DoublyLinkedList) Append(value int) {
	newNode := &Node{Value: value}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = newNode
}

// Prepend adds a new node with the given value to the front of the list
func (dll *DoublyLinkedList) Prepend(value int) {
	newNode := &Node{Value: value}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	newNode.Next = dll.Head
	dll.Head.Prev = newNode
	dll.Head = newNode
}

// DisplayForward prints the list from Head to Tail
func (dll *DoublyLinkedList) DisplayForward() {
	current := dll.Head
	for current != nil {
		log.Printf("%d -> ", current.Value)
		current = current.Next
	}
	log.Println("nil")
}

// DisplayBackward prints the list from Tail to Head
func (dll *DoublyLinkedList) DisplayBackward() {
	current := dll.Tail
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Prev
	}
	fmt.Println("nil")
}
