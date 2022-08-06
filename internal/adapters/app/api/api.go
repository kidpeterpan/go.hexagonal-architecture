package api

import "go.hexagonal-architecture/internal/ports"

type Adapter struct {
	arith ports.Arithmetic
	db    ports.DbPort
}

func NewAdapter(arith ports.Arithmetic, db ports.DbPort) *Adapter {
	return &Adapter{
		arith: arith,
		db:    db,
	}
}

func (a Adapter) GetAddition(x, y int32) (int32, error) {
	answer, err := a.arith.Additional(x, y)
	if err != nil {
		return 0, err
	}

	if err := a.db.AddToHistory(answer, "addition"); err != nil {
		return 0, err
	}
	return answer, err
}

func (a Adapter) GetSubtraction(x, y int32) (int32, error) {
	answer, err := a.arith.Subtraction(x, y)
	if err != nil {
		return 0, err
	}
	if err := a.db.AddToHistory(answer, "subtraction"); err != nil {
		return 0, err
	}
	return answer, err
}

func (a Adapter) GetMultiplication(x, y int32) (int32, error) {
	answer, err := a.arith.Multiplication(x, y)
	if err != nil {
		return 0, err
	}

	if err := a.db.AddToHistory(answer, "multiplication"); err != nil {
		return 0, err
	}
	return answer, err
}

func (a Adapter) GetDivision(x, y int32) (int32, error) {
	answer, err := a.arith.Division(x, y)
	if err != nil {
		return 0, err
	}

	if err := a.db.AddToHistory(answer, "division"); err != nil {
		return 0, err
	}

	return answer, err
}
