package oneof

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/bbengfort/oneof/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
)

const (
	bufsize   = 1024 * 1024
	buftarget = "bufnet"
)

const NMessages = 1000

type Server struct {
	pb.UnimplementedStreamerServer
	srv   *grpc.Server
	sock  *bufconn.Listener
	echan chan error
}

func New() (*Server, error) {
	s := &Server{
		srv:   grpc.NewServer(grpc.Creds(insecure.NewCredentials())),
		echan: make(chan error, 1),
		sock:  bufconn.Listen(bufsize),
	}
	pb.RegisterStreamerServer(s.srv, s)
	return s, nil
}

func (s *Server) Serve() (err error) {
	go func() {
		if err := s.srv.Serve(s.sock); err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				s.echan <- nil
			} else {
				s.echan <- err
			}
		}
	}()

	return <-s.echan
}

func (s *Server) Shutdown() {
	s.srv.GracefulStop()
	s.sock.Close()
}

func (s *Server) Onef(stream pb.Streamer_OnefServer) error {
	var (
		clientID   string
		region     string
		totalBytes int64
		acks       int64
		nacks      int64
		messages   int64
	)

	for i := 0; i < NMessages; i++ {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		messages++
		switch t := msg.Embed.(type) {
		case *pb.OnefMessage_Init:
			clientID = t.Init.ClientId
			region = t.Init.Region
		case *pb.OnefMessage_Data:
			totalBytes += int64(len(t.Data.Data))
		case *pb.OnefMessage_Ack:
			acks++
		case *pb.OnefMessage_Nack:
			nacks++
		}
	}

	// Send done message
	return stream.SendAndClose(&pb.Done{
		ServerId:   fmt.Sprintf("%s:%s", clientID, region),
		TotalBytes: totalBytes,
		Messages:   messages,
		Acks:       acks,
		Nacks:      nacks,
	})
}

func (s *Server) Anyf(stream pb.Streamer_AnyfServer) error {
	var (
		clientID   string
		region     string
		totalBytes int64
		acks       int64
		nacks      int64
		messages   int64
	)

	for i := 0; i < NMessages; i++ {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		messages++
		embed, err := msg.Embed.UnmarshalNew()
		if err != nil {
			return err
		}

		switch t := embed.(type) {
		case *pb.Init:
			clientID = t.ClientId
			region = t.Region
		case *pb.Data:
			totalBytes += int64(len(t.Data))
		case *pb.Ack:
			acks++
		case *pb.Nack:
			nacks++
		}
	}

	// Send done message
	return stream.SendAndClose(&pb.Done{
		ServerId:   fmt.Sprintf("%s:%s", clientID, region),
		TotalBytes: totalBytes,
		Messages:   messages,
		Acks:       acks,
		Nacks:      nacks,
	})
}

func (s *Server) Enum(stream pb.Streamer_EnumServer) error {
	var (
		clientID   string
		region     string
		totalBytes int64
		acks       int64
		nacks      int64
		messages   int64
	)

	for i := 0; i < NMessages; i++ {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		messages++
		switch msg.Mtype {
		case pb.EnumMessage_INIT:
			t := &pb.Init{}
			if err = proto.Unmarshal(msg.Embed, t); err != nil {
				return err
			}
			clientID = t.ClientId
			region = t.Region
		case pb.EnumMessage_DATA:
			t := &pb.Data{}
			if err = proto.Unmarshal(msg.Embed, t); err != nil {
				return err
			}
			totalBytes += int64(len(t.Data))
		case pb.EnumMessage_ACK:
			t := &pb.Ack{}
			if err = proto.Unmarshal(msg.Embed, t); err != nil {
				return err
			}
			acks++
		case pb.EnumMessage_NACK:
			t := &pb.Nack{}
			if err = proto.Unmarshal(msg.Embed, t); err != nil {
				return err
			}
			nacks++
		}
	}

	// Send done message
	return stream.SendAndClose(&pb.Done{
		ServerId:   fmt.Sprintf("%s:%s", clientID, region),
		TotalBytes: totalBytes,
		Messages:   messages,
		Acks:       acks,
		Nacks:      nacks,
	})
}

func (s *Server) Client(ctx context.Context) (pb.StreamerClient, error) {
	cc, err := grpc.DialContext(ctx, buftarget, grpc.WithContextDialer(s.Dial), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewStreamerClient(cc), nil
}

func (s *Server) Dial(context.Context, string) (net.Conn, error) {
	return s.sock.Dial()
}
