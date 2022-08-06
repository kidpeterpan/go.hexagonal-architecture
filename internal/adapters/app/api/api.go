package api

import "go.hexagonal-architecture/internal/ports"

type Adapter struct {
	arith ports.Arithmetic
}

func NewAdapter(arith ports.Arithmetic) *Adapter {
	return &Adapter{
		arith: arith,
	}
}

func (a Adapter) GetAddition(x, y int32) (int32, error) {
	answer, err := a.GetAddition(x, y)
	if err != nil {
		return 0, err
	}
	return answer, err
}

func (a Adapter) GetSubtraction(x, y int32) (int32, error) {
	answer, err := a.GetSubtraction(x, y)
	if err != nil {
		return 0, err
	}
	return answer, err
}

func (a Adapter) GetMultiplication(x, y int32) (int32, error) {
	answer, err := a.GetMultiplication(x, y)
	if err != nil {
		return 0, err
	}
	return answer, err
}

func (a Adapter) GetDivision(x, y int32) (int32, error) {
	answer, err := a.GetDivision(x, y)
	if err != nil {
		return 0, err
	}
	return answer, err
}
