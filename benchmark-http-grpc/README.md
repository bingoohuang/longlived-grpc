# Performance Benchmarking: gRPC+Protobuf vs. HTTP+JSON

![](snapshots/2024-08-04-22-20-26.png)

[Read the full article on packagemain.tech](https://packagemain.tech/p/protobuf-grpc-vs-json-http)

1. Generate unimplemented server gRPC client stub: `make pb`
2. Run benchmarks: `make bench`

```sh
$ make bench
go test -bench=. -benchmem=1  -benchtime=30s
goos: darwin
goarch: amd64
pkg: github.com/plutov/packagemain/benchmark-http-grpc
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkGRPCProtobuf-12               1        30003326104 ns/op         398136 B/op       1949 allocs/op
BenchmarkHTTP1JSON-12             466153             77131 ns/op            9243 B/op        108 allocs/op
BenchmarkHTTP2JSON-12             301926            126483 ns/op           12084 B/op        117 allocs/op
PASS
ok      github.com/plutov/packagemain/benchmark-http-grpc       107.851s

$  go version                                                     
go version go1.22.5 darwin/amd64

$ hyfetch
                                  bingoo@192
                    c.'           ----------
                 ,xNMM.           OS: macOS Sonoma 14.5 (23F79) x86_64
               .OMMMMo            Host: MacBook Pro (15-inch, 2018/2019)
               lMM"               Kernel: 23.5.0
     .;loddo:.  .olloddol;.       Uptime: 2 days, 3 hours, 32 mins
   cKMMMMMMMMMMNWMMMMMMMMMM0:     Packages: 5 (cargo), 1 (npm), 161 (brew)
 .KMMMMMMMMMMMMMMMMMMMMMMMWd.     Shell: zsh 5.9
 XMMMMMMMMMMMMMMMMMMMMMMMX.       Resolution: 1920x1200 @ 2x
;MMMMMMMMMMMMMMMMMMMMMMMM:        DE: Aqua
:MMMMMMMMMMMMMMMMMMMMMMMM:        WM: Rectangle
.MMMMMMMMMMMMMMMMMMMMMMMMX.       Terminal: iTerm2
 kMMMMMMMMMMMMMMMMMMMMMMMMWd.     Terminal Font: MesloLGS-NF-Regular 13
 'XMMMMMMMMMMMMMMMMMMMMMMMMMMk    CPU: Intel i7-8750H (12) @ 2.20GHz
  'XMMMMMMMMMMMMMMMMMMMMMMMMK.    GPU: Intel UHD Graphics 630, Radeon Pro 555X
    kMMMMMMMMMMMMMMMMMMMMMMd      Memory: 11.13 GiB / 16.00 GiB (69%)
     ;KMMMMMMMWXXWMMMMMMMk.       Network: en0: Wi-Fi@Mbps
       "cooc*"    "*coo'"
```

- [hyfetch](https://github.com/hykilpikonna/hyfetch)
