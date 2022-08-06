package arithmetic

type Arithmetic struct {
}

func NewArithmetic() *Arithmetic {
	return &Arithmetic{}
}

func (a Arithmetic) Additional(x int32, y int32) (int32, error) {
	return x + y, nil
}

func (a Arithmetic) Subtraction(x int32, y int32) (int32, error) {
	return x - y, nil
}

func (a Arithmetic) Multiplication(x int32, y int32) (int32, error) {
	return x * y, nil
}
func (a Arithmetic) Division(x int32, y int32) (int32, error) {
	return x / y, nil
}
