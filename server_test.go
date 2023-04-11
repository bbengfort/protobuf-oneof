package oneof_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/bbengfort/oneof"
	"github.com/bbengfort/oneof/pb"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestServer(t *testing.T) {
	srv, err := oneof.New()
	require.NoError(t, err)

	go srv.Serve()
	t.Cleanup(srv.Shutdown)

	client, err := srv.Client(context.Background())
	require.NoError(t, err)

	messages, totalBytes, acks, nacks := makeMessages()

	t.Run("Onef", func(t *testing.T) {
		stream, err := client.Onef(context.Background())
		require.NoError(t, err)

		for _, msg := range messages {
			var onefm *pb.OnefMessage
			switch t := msg.(type) {
			case *pb.Init:
				onefm = &pb.OnefMessage{
					Embed: &pb.OnefMessage_Init{
						Init: t,
					},
				}
			case *pb.Data:
				onefm = &pb.OnefMessage{
					Embed: &pb.OnefMessage_Data{
						Data: t,
					},
				}
			case *pb.Ack:
				onefm = &pb.OnefMessage{
					Embed: &pb.OnefMessage_Ack{
						Ack: t,
					},
				}
			case *pb.Nack:
				onefm = &pb.OnefMessage{
					Embed: &pb.OnefMessage_Nack{
						Nack: t,
					},
				}
			}

			err := stream.Send(onefm)
			require.NoError(t, err)
		}

		done, err := stream.CloseAndRecv()
		require.NoError(t, err)
		require.Equal(t, "testing:localhost", done.ServerId)
		require.Equal(t, int64(oneof.NMessages), done.Messages)
		require.Equal(t, totalBytes, done.TotalBytes)
		require.Equal(t, acks, done.Acks)
		require.Equal(t, nacks, done.Nacks)
	})

	t.Run("Anyf", func(t *testing.T) {
		stream, err := client.Anyf(context.Background())
		require.NoError(t, err)

		for _, msg := range messages {
			anym := &pb.AnyfMessage{}
			anym.Embed, err = anypb.New(msg)
			require.NoError(t, err)

			err = stream.Send(anym)
			require.NoError(t, err)
		}

		done, err := stream.CloseAndRecv()
		require.NoError(t, err)
		require.Equal(t, "testing:localhost", done.ServerId)
		require.Equal(t, int64(oneof.NMessages), done.Messages)
		require.Equal(t, totalBytes, done.TotalBytes)
		require.Equal(t, acks, done.Acks)
		require.Equal(t, nacks, done.Nacks)
	})

	t.Run("Enum", func(t *testing.T) {
		stream, err := client.Enum(context.Background())
		require.NoError(t, err)

		for _, msg := range messages {
			var enumm *pb.EnumMessage
			switch msg.(type) {
			case *pb.Init:
				enumm = &pb.EnumMessage{
					Mtype: pb.EnumMessage_INIT,
				}
			case *pb.Data:
				enumm = &pb.EnumMessage{
					Mtype: pb.EnumMessage_DATA,
				}
			case *pb.Ack:
				enumm = &pb.EnumMessage{
					Mtype: pb.EnumMessage_ACK,
				}
			case *pb.Nack:
				enumm = &pb.EnumMessage{
					Mtype: pb.EnumMessage_NACK,
				}
			}

			enumm.Embed, err = proto.Marshal(msg)
			require.NoError(t, err)

			err := stream.Send(enumm)
			require.NoError(t, err)
		}

		done, err := stream.CloseAndRecv()
		require.NoError(t, err)
		require.Equal(t, "testing:localhost", done.ServerId)
		require.Equal(t, int64(oneof.NMessages), done.Messages)
		require.Equal(t, totalBytes, done.TotalBytes)
		require.Equal(t, acks, done.Acks)
		require.Equal(t, nacks, done.Nacks)
	})

}

func BenchmarkServer(b *testing.B) {
	srv, err := oneof.New()
	require.NoError(b, err)

	go srv.Serve()
	b.Cleanup(srv.Shutdown)

	client, err := srv.Client(context.Background())
	require.NoError(b, err)

	messages, _, _, _ := makeMessages()

	b.Run("Onef", func(b *testing.B) {
		stream, err := client.Onef(context.Background())
		require.NoError(b, err)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {

			for _, msg := range messages {
				var onefm *pb.OnefMessage
				switch t := msg.(type) {
				case *pb.Init:
					onefm = &pb.OnefMessage{
						Embed: &pb.OnefMessage_Init{
							Init: t,
						},
					}
				case *pb.Data:
					onefm = &pb.OnefMessage{
						Embed: &pb.OnefMessage_Data{
							Data: t,
						},
					}
				case *pb.Ack:
					onefm = &pb.OnefMessage{
						Embed: &pb.OnefMessage_Ack{
							Ack: t,
						},
					}
				case *pb.Nack:
					onefm = &pb.OnefMessage{
						Embed: &pb.OnefMessage_Nack{
							Nack: t,
						},
					}
				}

				stream.Send(onefm)
			}
		}
	})

	b.Run("Anyf", func(b *testing.B) {
		stream, err := client.Anyf(context.Background())
		require.NoError(b, err)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, msg := range messages {
				anym := &pb.AnyfMessage{}
				anym.Embed, err = anypb.New(msg)
				require.NoError(b, err)

				stream.Send(anym)
			}
		}
	})

	b.Run("Enum", func(b *testing.B) {
		stream, err := client.Enum(context.Background())
		require.NoError(b, err)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, msg := range messages {
				var enumm *pb.EnumMessage
				switch msg.(type) {
				case *pb.Init:
					enumm = &pb.EnumMessage{
						Mtype: pb.EnumMessage_INIT,
					}
				case *pb.Data:
					enumm = &pb.EnumMessage{
						Mtype: pb.EnumMessage_DATA,
					}
				case *pb.Ack:
					enumm = &pb.EnumMessage{
						Mtype: pb.EnumMessage_ACK,
					}
				case *pb.Nack:
					enumm = &pb.EnumMessage{
						Mtype: pb.EnumMessage_NACK,
					}
				}

				enumm.Embed, err = proto.Marshal(msg)
				require.NoError(b, err)

				stream.Send(enumm)
			}
		}
	})

}

var result []byte

func BenchmarkMarshal(b *testing.B) {
	b.Run("Oneof", func(b *testing.B) {
		var d []byte
		var err error

		b.Run("FixedSize", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				embed := makeNack()
				b.StartTimer()

				msg := &pb.OnefMessage{
					Embed: &pb.OnefMessage_Nack{
						Nack: embed,
					},
				}

				d, err = proto.Marshal(msg)
				require.NoError(b, err)

				omsg := &pb.OnefMessage{}
				err = proto.Unmarshal(d, omsg)
				require.NoError(b, err)

				b.SetBytes(int64(len(d)))
			}
			result = d
		})

		b.Run("VariableSize", func(b *testing.B) {
			var d []byte
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				embed := makeData()
				b.StartTimer()

				msg := &pb.OnefMessage{
					Embed: &pb.OnefMessage_Data{
						Data: embed,
					},
				}

				d, err = proto.Marshal(msg)
				require.NoError(b, err)

				omsg := &pb.OnefMessage{}
				err = proto.Unmarshal(d, omsg)
				require.NoError(b, err)

				b.SetBytes(int64(len(d)))
			}
			result = d
		})
	})

	b.Run("Anyf", func(b *testing.B) {
		var d []byte

		b.Run("FixedSize", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				embed := makeNack()
				b.StartTimer()

				msg, err := anypb.New(embed)
				require.NoError(b, err)

				d, err = proto.Marshal(msg)
				require.NoError(b, err)

				omsg := &anypb.Any{}
				err = proto.Unmarshal(d, omsg)
				require.NoError(b, err)

				nack := &pb.Nack{}
				err = omsg.UnmarshalTo(nack)
				require.NoError(b, err)

				b.SetBytes(int64(len(d)))
			}
			result = d
		})

		b.Run("VariableSize", func(b *testing.B) {
			var d []byte
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				embed := makeData()
				b.StartTimer()

				msg, err := anypb.New(embed)
				require.NoError(b, err)

				d, err = proto.Marshal(msg)
				require.NoError(b, err)

				omsg := &anypb.Any{}
				err = proto.Unmarshal(d, omsg)
				require.NoError(b, err)

				data := &pb.Data{}
				err = omsg.UnmarshalTo(data)
				require.NoError(b, err)

				b.SetBytes(int64(len(d)))
			}
			result = d
		})
	})

	b.Run("Enum", func(b *testing.B) {
		var d []byte
		var err error

		b.Run("FixedSize", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				embed := makeNack()
				msg := &pb.EnumMessage{
					Mtype: pb.EnumMessage_NACK,
				}
				b.StartTimer()

				msg.Embed, err = proto.Marshal(embed)
				require.NoError(b, err)

				d, err = proto.Marshal(msg)
				require.NoError(b, err)

				omsg := &pb.EnumMessage{}
				err = proto.Unmarshal(d, omsg)
				require.NoError(b, err)

				nack := &pb.Nack{}
				err = proto.Unmarshal(omsg.Embed, nack)
				require.NoError(b, err)

				b.SetBytes(int64(len(d)))
			}
			result = d
		})

		b.Run("VariableSize", func(b *testing.B) {
			var d []byte
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				embed := makeData()
				msg := &pb.EnumMessage{
					Mtype: pb.EnumMessage_NACK,
				}
				b.StartTimer()

				msg.Embed, err = proto.Marshal(embed)
				require.NoError(b, err)

				d, err = proto.Marshal(msg)
				require.NoError(b, err)

				omsg := &pb.EnumMessage{}
				err = proto.Unmarshal(d, omsg)
				require.NoError(b, err)

				data := &pb.Data{}
				err = proto.Unmarshal(omsg.Embed, data)
				require.NoError(b, err)

				b.SetBytes(int64(len(d)))
			}
			result = d
		})
	})
}

func makeNack() *pb.Nack {
	return &pb.Nack{
		Id:    ulid.Make().String(),
		Code:  uint32(rand.Intn(600) + 100),
		Error: "this is a random error",
	}
}

func makeData() *pb.Data {
	size := rand.Intn(1024*1024) + 50
	dmsg := &pb.Data{
		Id:   ulid.Make().String(),
		Data: make([]byte, size),
	}
	rand.Read(dmsg.Data)
	return dmsg
}

func makeMessages() ([]proto.Message, int64, int64, int64) {
	messages := make([]proto.Message, 0, oneof.NMessages)
	messages = append(messages, &pb.Init{ClientId: "testing", Region: "localhost"})

	var acks, nacks, totalBytes int64
	for i := 0; i < oneof.NMessages-1; i++ {
		var msg proto.Message
		if i%2 == 0 {
			// Create a data message
			dmsg := makeData()
			totalBytes += int64(len(dmsg.Data))
			msg = dmsg
		} else {
			// Create an ack or a nack with some probability
			coinFlip := rand.Float64()
			if coinFlip < 0.1 {
				nacks++
				msg = makeNack()
			} else {
				acks++
				msg = &pb.Ack{
					Id: ulid.Make().String(),
					Ts: timestamppb.Now(),
				}
			}
		}

		messages = append(messages, msg)
	}
	return messages, totalBytes, acks, nacks
}
