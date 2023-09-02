package request

type StuRequest struct {
	Rev       string `json:"_rev,omitempty"`
	StuName   string `json:"stuname"`
	Gender    string `json:"gender"`
	ClassName string `json:"classname"`
	Note      string `json:"note"`
}
