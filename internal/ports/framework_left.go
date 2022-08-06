package ports

type GrpcPort interface {
	Run()
	GetAddition()
	GetSubtraction()
	GetMultiplication()
	GetDivision()
}
