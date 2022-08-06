package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a Adapter) Additional(x int32, y int32) (int32, error) {
	return x + y, nil
}

func (a Adapter) Subtraction(x int32, y int32) (int32, error) {
	return x - y, nil
}

func (a Adapter) Multiplication(x int32, y int32) (int32, error) {
	return x * y, nil
}
func (a Adapter) Division(x int32, y int32) (int32, error) {
	return x / y, nil
}
