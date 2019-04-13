package okex

type Currency struct {
	// api/account/v3/currencies

	Name          string  `json:"name"`
	Currency      string  `json:"currency"`
	CanWithdraw   int     `json:"can_withdraw"`
	CanDeposit    int     `json:"can_deposit"`
	MinWithdrawal float64 `json:"min_withdrawal"`
}
