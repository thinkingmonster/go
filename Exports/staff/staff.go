package staff

var OverPaidLimit = 50000
var UnderPaidLimit = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (c Office) All() []Employee {
	return c.AllStaff
}

func (c Office) OverPaid() []Employee {
	var Overpaid []Employee
	for _, x := range c.AllStaff {
		if x.Salary > OverPaidLimit {
			Overpaid = append(Overpaid, x)
		}
	}
	return Overpaid
}

func (c Office) UnderPaid() []Employee {
	var UnderPaid []Employee
	for _, x := range c.AllStaff {
		if x.Salary < UnderPaidLimit {
			UnderPaid = append(UnderPaid, x)
		}
	}
	return UnderPaid
}
