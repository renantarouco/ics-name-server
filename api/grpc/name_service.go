package grpc

import (
	"io"

	"github.com/renantarouco/ics-name-server/api/grpc/proto"
)

// NameServiceServer - Struct holding the NameServiceServer implementation
type NameServiceServer struct{}

// ConnectMessageServer - Connect function used by the client
func (s *NameServiceServer) ConnectMessageServer(stream proto.NameService_ConnectMessageServerServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		connectionInfo := &proto.ConnectionInfo{
			Header: &proto.ConnectionHeader{
				MessageServerId: "new",
			},
		}
		if err := stream.Send(connectionInfo); err != nil {
			return err
		}
	}
}
