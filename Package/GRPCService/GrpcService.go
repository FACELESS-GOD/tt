package GrpcService

import (
	"context"

	tt_git "github.com/FACELESS-GOD/tt.git/Package/ProtoSystemGenerated/PROTOBUF"
)

type GrpcService struct {
	tt_git.UnimplementedTestGRPCServiceServer
}

func (grpcService *GrpcService) GetData(Req *tt_git.DataReceipt, Srv tt_git.TestGRPCService_GetDataServer) error {
	items := []*tt_git.Item{
		&tt_git.Item{
			Id:   "1",
			Name: "a",
		},
		&tt_git.Item{
			Id:   "2",
			Name: "b",
		},
		&tt_git.Item{
			Id:   "3",
			Name: "c",
		},
	}

	for i, _ := range items {
		Srv.Send(
			&tt_git.Data{
				Items: items[0 : i+1],
			})
	}

	return nil

}

func (grpcService *GrpcService) AddData(Ctx context.Context, Req *tt_git.SendData) (Data *tt_git.DataReceipt, err error) {
	return &tt_git.DataReceipt{
		ID: "10",
	}, nil
}

func (grpcService *GrpcService) GetDataStatus(Ctx context.Context, Req *tt_git.DataReceipt) (Data *tt_git.DataStatus, err error) {

	return &tt_git.DataStatus{
		DataID: Req.ID,
		Status: "OK",
	}, nil

}
