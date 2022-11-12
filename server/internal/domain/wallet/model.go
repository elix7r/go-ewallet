package wallet

type Wallet struct {
	UUID    string  `json:"_id,omitempty"`
	Balance float64 `json:"balance"`
}
