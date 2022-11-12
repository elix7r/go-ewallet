package transaction

type Type int64

const (
	Transfer Type = iota
)

type Transaction struct {
	UUID     string  `json:"_id,omitempty"`
	Type     Type    `json:"type,omitempty"`
	From     string  `json:"from,omitempty"`
	To       string  `json:"to,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Date     string  `json:"date,omitempty"`
	IsIssued bool    `json:"is_issued"`
}
