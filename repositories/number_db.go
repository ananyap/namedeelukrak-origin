package repositories

import (
	"github.com/jmoiron/sqlx"
)

type NumberRepositoryDB struct {
	db *sqlx.DB
}

func NewRepositoryDB(db *sqlx.DB) NumberRepository {
	return NumberRepositoryDB{db}
}

func (r NumberRepositoryDB) GetAll() ([]Number, error) {
	numbers := []Number{}
	query := "SELECT pairnumberid, pairnumber, pairtype, pairpoint, miracledesc FROM numbers"
	err := r.db.Select(&numbers, query)
	if err != nil {
		return nil, err
	}
	return numbers, nil
}
func (r NumberRepositoryDB) GetById(id int) (*Number, error) {
	number := Number{}
	query := "SELECT pairnumberid, pairnumber, pairtype, pairpoint, miracledesc FROM numbers WHERE pairnumberid = ?"
	err := r.db.Get(&number, query, id)
	if err != nil {
		return nil, err
	}
	return &number, nil
}
