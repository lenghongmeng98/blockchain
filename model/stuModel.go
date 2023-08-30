package model

type Stu struct {
	ID        string `json:"_id"`
	Rev       string `json:"_rev,omitempty"`
	StuName   string `json:"stuname"`
	Gender    string `json:"gender"`
	ClassName string `json:"classname"`
	Note      string `json:"note"`
}
