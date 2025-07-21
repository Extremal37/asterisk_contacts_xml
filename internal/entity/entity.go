package entity

type ContactDB struct {
	number string `db:"name"`
	name   string `db:"callerid"`
}
