package visual_novel

type ScriptEntry struct {
	Type        string `json:"type"`
	Action      string `json:"action"`
	ID          int    `json:"id"`
	Label       string `json:"label,omitempty"`
	Sprite      string `json:"sprite,omitempty"`
	Location    string `json:"location,omitempty"`
	DynLocation string `json:"dyn_location,omitempty"`
	Position    string `json:"position,omitempty"`
	WRatio      int    `json:"wRatio,omitempty"`
	HRatio      int    `json:"hRatio,omitempty"`
	WFrameRatio int    `json:"wFrameRatio,omitempty"`
	HFrameRatio int    `json:"hFrameRatio,omitempty"`
	Column      int    `json:"column,omitempty"`
	Row         int    `json:"row,omitempty"`
	Content     string `json:"content,omitempty"`
	Actions     string `json:"actions,omitempty"`
}
