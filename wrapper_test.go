package oneof_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/bbengfort/oneof"
	"github.com/bbengfort/oneof/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func BenchmarkWrapper(b *testing.B) {
	// b.Run("Tiny", makeBenchmark(1024))
	// b.Run("XSmall", makeBenchmark(4096))
	// b.Run("Small", makeBenchmark(8192))
	// b.Run("Medium", makeBenchmark(32768))
	// b.Run("Large", makeBenchmark(1.468e+6))
	// b.Run("XLarge", makeBenchmark(5.243e6))

	for i := 3; i < 23; i++ {
		s := int(math.Pow(2, float64(i)))
		b.Run(fmt.Sprintf("%d", s), makeBenchmark(s))
	}

}

var resultmsg proto.Message

func makeBenchmark(size int) func(b *testing.B) {
	return func(b *testing.B) {

		flat := oneof.Flat(size)
		wrapped := oneof.FlatToWrapped(flat)

		flatData, err := proto.Marshal(flat)
		require.NoError(b, err)

		wrappedData, err := proto.Marshal(wrapped)
		require.NoError(b, err)

		b.Run("FlatMarshal", func(b *testing.B) {
			var data []byte
			for i := 0; i < b.N; i++ {
				data, _ = proto.Marshal(flat)
			}
			result = data
		})

		b.Run("WrappedMarshal", func(b *testing.B) {
			var data []byte
			for i := 0; i < b.N; i++ {
				data, _ = proto.Marshal(wrapped)
			}
			result = data
		})

		b.Run("FlatUnmarshal", func(b *testing.B) {
			var msg *pb.Flat
			for i := 0; i < b.N; i++ {
				msg = &pb.Flat{}
				proto.Unmarshal(flatData, msg)
			}
			resultmsg = msg
		})

		b.Run("WrappedUnmarshal", func(b *testing.B) {
			var msg *pb.Wrapped
			for i := 0; i < b.N; i++ {
				msg = &pb.Wrapped{}
				proto.Unmarshal(wrappedData, msg)
			}
			resultmsg = msg
		})

	}
}
