package packet

type Content struct {
	Color  string      `json:"color,omitempty"`
	Size   int         `json:"size,omitempty"`
	Style  string      `json:"style,omitempty"`
	Action interface{} `json:"action,omitempty"`

	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	Img  string `json:"img,omitempty"`
}
