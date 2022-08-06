package ports

type ArithmeticPort interface {
	Additional(x int32, y int32) (int32, error)
	Subtraction(x int32, y int32) (int32, error)
	Multiplication(x int32, y int32) (int32, error)
	Division(x int32, y int32) (int32, error)
}
