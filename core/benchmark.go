package core

var inf Info = Make_Info(
	"https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-11.7.0-amd64-netinst.iso",
	"debian-11.7.0-amd64-netinst.iso",
	4,
)

func Benchmark_Seq() {
	Download_Seq(inf)
}

func Benchmark_Par() {
	Download(inf)
}
