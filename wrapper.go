package oneof

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"strings"

	"github.com/bbengfort/oneof/pb"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Flat(size int) *pb.Flat {
	return &pb.Flat{
		Id:            ulid.Make().String(),
		TopicId:       ulid.Make().String(),
		Mimetype:      RandEnum(90),
		Type:          RandType(),
		Key:           RandBytes(20),
		Data:          RandBytes(size),
		Encryption:    RandEncryption(),
		Compression:   RandCompression(),
		Geography:     RandRegion(),
		Publisher:     RandPublisher(),
		UserDefinedId: AlphaNumeric(14),
		Created:       timestamppb.Now(),
		Committed:     timestamppb.Now(),
	}
}

func FlatToWrapped(flat *pb.Flat) *pb.Wrapped {
	inner := &pb.Inner{
		Mimetype:      flat.Mimetype,
		Type:          flat.Type,
		Encryption:    flat.Encryption,
		Compression:   flat.Compression,
		UserDefinedId: flat.UserDefinedId,
		Data:          flat.Data,
		Created:       flat.Created,
	}

	wrapped := &pb.Wrapped{
		Id:        flat.Id,
		TopicId:   flat.TopicId,
		Geography: flat.Geography,
		Publisher: flat.Publisher,
		Key:       flat.Key,
		Committed: flat.Committed,
	}

	var err error
	if wrapped.Inner, err = proto.Marshal(inner); err != nil {
		panic(err)
	}
	return wrapped
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	idxbits  = 6
	idxmask  = 1<<idxbits - 1
	idxmax   = 63 / idxbits
)

// Alpha generates a random string of n characters that only includes upper and
// lowercase letters (no symbols or digits).
func Alpha(n int) string {
	return generate(n, alphabet)
}

// AlphaNumeric generates a random string of n characters that includes upper and
// lowercase letters and the digits 0-9.
func AlphaNumeric(n int) string {
	return generate(n, alphanum)
}

// generate is a helper function to create a random string of n characters from the
// character set defined by chars. It uses as efficient a method of generation as
// possible, using a string builder to prevent multiple allocations and a 6 bit mask
// to select 10 random letters at a time to add to the string. This method would be far
// faster if it used math/rand src and the Int63() function, but for API key generation
// it is important to use a cryptographically random generator.
//
// See: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func generate(n int, chars string) string {
	sb := strings.Builder{}
	sb.Grow(n)

	for i, cache, remain := n-1, rand.Int63(), idxmax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), idxmax
		}

		if idx := int(cache & idxmask); idx < len(chars) {
			sb.WriteByte(chars[idx])
			i--
		}

		cache >>= idxbits
		remain--
	}

	return sb.String()
}

func RandEnum(max int64) uint32 {
	e := rand.Int63n(max)
	return uint32(e)
}

func RandBytes(n int) []byte {
	data := make([]byte, n)
	if _, err := crand.Read(data); err != nil {
		panic(err)
	}
	return data
}

func RandIPAddr() string {
	octs := make([]any, 4)
	for i := 0; i < 4; i++ {
		octs[i] = rand.Int63n(255) + 1
	}
	return fmt.Sprintf("%d.%d.%d.%d", octs...)
}

func RandType() *pb.Type {
	return &pb.Type{
		Name:         Alpha(8),
		MajorVersion: RandEnum(6),
		MinorVersion: RandEnum(44),
		PatchVersion: RandEnum(10),
	}
}

func RandEncryption() *pb.Encryption {
	return &pb.Encryption{
		Algorithm:     RandEnum(11),
		KeyId:         Alpha(32),
		EncryptionKey: RandBytes(32),
		HmacSecret:    RandBytes(32),
		Signature:     RandBytes(128),
	}
}

func RandCompression() *pb.Compression {
	return &pb.Compression{
		Algorithm: RandEnum(7),
	}
}

func RandRegion() *pb.Region {
	country := Alpha(2)
	return &pb.Region{
		Region:  fmt.Sprintf("%s-%s", country, Alpha(9)),
		Zone:    Alpha(1),
		Country: country,
	}
}

func RandPublisher() *pb.Publisher {
	return &pb.Publisher{
		PublisherId: ulid.Make().String(),
		Ipaddr:      RandIPAddr(),
		ClientId:    Alpha(12),
		UserAgent:   Alpha(9),
	}
}
