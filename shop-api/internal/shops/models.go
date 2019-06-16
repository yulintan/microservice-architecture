package shops

type Shop struct {
	ID         int    `json:"id" db:"id"`
	ShopDomain string `json:"shop_domain" db:"shop_domain"`
	Currency   string `json:"currency" db:"currency"`
}
