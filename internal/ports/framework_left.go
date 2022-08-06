package ports

// this not work because of can not implement grpc outside pb package
// reason the interface contains mustEmbedUnimplementedArithmeticServiceServer()

// type GrpcPort interface {
//	   Run()
//	   GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
//	   GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
//	   GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
//	   GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
// }
