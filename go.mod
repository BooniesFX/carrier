module github.com/IPFS-eX/carrier

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.0
	github.com/huin/goupnp v1.0.0
	github.com/jackpal/go-nat-pmp v1.0.2
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/klauspost/reedsolomon v1.9.4 // indirect
	github.com/lucas-clemente/quic-go v0.15.4
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/pkg/errors v0.9.1
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.5.1
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20191217153810-f85b25db303b // indirect
	github.com/tjfoc/gmsm v1.3.0 // indirect
	github.com/xtaci/kcp-go v5.4.20+incompatible
	github.com/xtaci/lossyconn v0.0.0-20200209145036-adba10fffc37 // indirect
	github.com/xtaci/smux v1.5.12
)

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.12.1

replace github.com/marten-seemann/qtls => github.com/marten-seemann/qtls v0.3.2
