package storage

type Category struct {
	ID 				int64 `db:"id"`
	CategoryName	string `db:"category_name"`
	IsCompleted		bool	`db:"is_completed"`
}

type Post struct {
	ID				int64 `db:"id"`
	CategoryID		int64 `db:"cat_id"`
	Title			string`db:"title"`
	Description		string`db:"description"`
	Image			string`db:"image"`
	CategoryName	string`db:"category_name"`
}