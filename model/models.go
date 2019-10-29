package model

// A Component is a trigger, target or effect of a trap
// does not work yet
type Component struct {
	ID   int    `json:"_id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
	Text string `json:"text"`
	Cost int    `json:"cost,omitempty"`
	P1   string `json:"param1,omitempty"`
	P2   string `json:"param2,omitempty"`
	P4   string `json:"param4,omitempty"`
	P3   string `json:"param3,omitempty"`
}

// An Upgrade is a trigger, target or effect of a trap
type Upgrade struct {
	ID          int    `json:"_id,omitempty"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Text        string `json:"text"`
	Cost        int    `json:"cost"`
	ComponentID int    `json:"component_id"`
}
