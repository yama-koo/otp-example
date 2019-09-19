package totp

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

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
	fmt.Printf("String:       %s\n", key.String())
	fmt.Printf("URL:          %s\n", key.URL())
	fmt.Println("Writing PNG to qr-code.png....")
	_ = ioutil.WriteFile("qr-code-totp.png", data, 0644)
	fmt.Println("")
	fmt.Println("Please add your TOTP to your OTP Application now!")
	fmt.Println("")
}

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

func Run() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
		// Algorithm:   otp.AlgorithmSHA256,
	})
	if err != nil {
		panic(err)
	}
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	_ = png.Encode(&buf, img)

	// display the QR code to the user.
	display(key, buf.Bytes())

	// Now Validate that the user's successfully added the passcode.
	fmt.Println("Validating TOTP...")
	passcode := promptForPasscode()
	valid := totp.Validate(passcode, key.Secret())
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passocde!")
		os.Exit(1)
	}
}
