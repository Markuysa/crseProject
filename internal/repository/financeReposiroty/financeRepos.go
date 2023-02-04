package financereposiroty

type Repository struct {
	moneyWaste map[string]int
}

type ExpenditureDataKeeper interface {
	AddNewExpenditure()
}

func (r *Repository) AddNewExpenditure(category string, amount int) error {
	r.moneyWaste[category]+=amount;
	return nil
}

func New() *Repository{
	return &Repository{}
}

