package expenditurereport


type Report struct{

	WastedMoney map[string]int
}
func New(WastedMoney map[string]int) *Report{

	return &Report{
		WastedMoney: WastedMoney,
	}
	
}

func (r *Report)GetReport(){

	
}