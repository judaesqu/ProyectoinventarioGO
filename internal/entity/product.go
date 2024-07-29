package entity

type Product struct {
	ID          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float32 `db:"price"`
	CreatedBy   int     `db:"created_by"`
}
