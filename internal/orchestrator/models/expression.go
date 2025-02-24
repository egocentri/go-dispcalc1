package models

type Expression struct {
	ID         int     `json:"id"`
	Expression string  `json:"expression"`
	Status     string  `json:"status"` 
	Result     float64 `json:"result,omitempty"`
}
