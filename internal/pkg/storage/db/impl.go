package db

func NewStorage(db DBTX) Storage {
	return &Queries{db}
}
