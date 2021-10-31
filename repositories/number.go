package repositories

type Number struct {
	PairNumberid int    `sql:"pairnumberid"`
	PairNumber   string `sql:"pairnumber"`
	PairType     string `sql:"pairtype"`
	PairPoint    int    `sql:"pairpoint"`
	MiracleDesc  string `sql:"miracledesc"`
}

type NumberRepository interface {
	GetAll() ([]Number, error)
	GetById(id int) (*Number, error)
}
