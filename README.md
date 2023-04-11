# ProtocolBuffer Oneof

The gRPC stream type is defined by a stream of a single type of messages to and from the server. However, we often want to have multiple message types used in the stream. This repository explores three methods for doing that:

1. A [`oneof`](https://protobuf.dev/programming-guides/proto3/#oneof) field
2. Using [`anypb`](https://protobuf.dev/programming-guides/proto3/#any)
3. Using an enum and marshaling/unmarshaling by hand

The benchmarks are as follows:

```
goos: darwin
goarch: arm64
pkg: github.com/bbengfort/oneof
BenchmarkServer/Onef-10         	      92	  12615174 ns/op	267043702 B/op	    6300 allocs/op
BenchmarkServer/Anyf-10         	      42	  26170065 ns/op	541325439 B/op	    8711 allocs/op
BenchmarkServer/Enum-10         	      43	  25941439 ns/op	540792656 B/op	    6647 allocs/op
```

```
BenchmarkMarshal/Oneof/FixedSize-10         	  575956	      2347 ns/op	  24.29 MB/s	     344 B/op	       8 allocs/op
BenchmarkMarshal/Anyf/FixedSize-10          	  385371	      2948 ns/op	  30.19 MB/s	     584 B/op	      10 allocs/op
BenchmarkMarshal/Enum/FixedSize-10          	  483910	      2443 ns/op	  24.15 MB/s	     408 B/op	       7 allocs/op
BenchmarkMarshal/Oneof/VariableSize-10      	   15384	     81611 ns/op	7163.58 MB/s	 1061618 B/op	       8 allocs/op
BenchmarkMarshal/Anyf/VariableSize-10       	    5978	    176417 ns/op	2090.65 MB/s	 2121335 B/op	      10 allocs/op
BenchmarkMarshal/Enum/VariableSize-10       	    6966	    175855 ns/op	3974.48 MB/s	 2140438 B/op	       7 allocs/op
```

Analysis will be provided in a future blog post.