package data

type Todo struct {
	ID          int64 `bun:",pk,autoincrement"`
	Title       string
	Description string
}
