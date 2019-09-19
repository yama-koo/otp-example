package hotp

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/hotp"

	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
)

func display(key *otp.Key, data []byte) {
	fmt.Printf("Issuer:       %s\n", key.Issuer())
	fmt.Printf("Account Name: %s\n", key.AccountName())
	fmt.Printf("Secret:       %s\n", key.Secret())
	fmt.Printf("Secret:       %s\n", key.Secret())
	fmt.Println("Writing PNG to qr-code.png....")
	_ = ioutil.WriteFile("qr-code-hotp.png", data, 0644)
	fmt.Println("")
	fmt.Println("Please add your HOTP to your OTP Application now!")
	fmt.Println("")
}

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

func Run() {
	key, err := hotp.Generate(hotp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
		// Algorithm:   otp.AlgorithmSHA256,
	})
	if err != nil {
		panic(err)
	}
	// Convert HOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	_ = png.Encode(&buf, img)

	// display the QR code to the user.
	display(key, buf.Bytes())

	// Now Validate that the user's successfully added the passcode.
	fmt.Println("Validating HOTP...")
	passcode := promptForPasscode()
	valid := hotp.Validate(passcode, 0, key.Secret())
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passocde!")
		os.Exit(1)
	}
}
