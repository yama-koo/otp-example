package main

import "github.com/yama-koo/otp-example/totp"

func main() {
	totp.Run()
	// var b32NoPadding = base32.StdEncoding.WithPadding(base32.NoPadding)

	// str := b32NoPadding.EncodeToString([]byte("hogeaaaaaaaaaaaaaaaa"))

	// str, _ := bcrypt.GenerateFromPassword([]byte("hogeaaaaaaaaaaaaaaaa"), bcrypt.DefaultCost)

	// fmt.Printf("%s\n", string(str))
}
