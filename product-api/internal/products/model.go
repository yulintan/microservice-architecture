package products

type Product struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Price    string `json:"price" db:"price"`
	Currency string `json:"currency" db:"currency"`
}
