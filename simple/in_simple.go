package simple_package

type parent struct {
	A int
}

func NewA() parent {
	return parent{10}
}