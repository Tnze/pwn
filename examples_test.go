package pwn_test

import "github.com/Tnze/pwn/v2"

func Example() {
	p := pwn.Remote("example.com:1314")
	// p:=pwn.Local(exec.Command("./a.out"))

	p.Write([]byte{0x00, 0x01, 0x02}) // payload

	p.Interactive()
}
