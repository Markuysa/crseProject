package financeservice

import (
	repos "ozonProjectmodule/internal/repository/financeReposiroty"
	"github.com/pkg/errors"
)

var repository = repos.New()

func Add(category string, amount int) error{
	
	err := repository.AddNewExpenditure(category,amount)
	if err!=nil{
		return errors.Wrap(err,"adding new expenditure")
	}
	return nil

}
