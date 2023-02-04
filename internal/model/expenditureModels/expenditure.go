package expendituremodels


import "time"

type Expenditure struct {
	Anmount         int64
	ExpenditureType string
	Date            time.Time 		
}
	
func New(amount int64, expenditureType string,
	date time.Time) *Expenditure{

		return &Expenditure{
			Anmount: amount,
			ExpenditureType: expenditureType,
			Date: date,
		}
		
	}