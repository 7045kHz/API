package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
CREDIT: https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
*/
type Authorization struct {
	Administrator User `json:"Administrator"`
}
type User struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func Help(e, s string) {
	fmt.Printf("This message is generated if %s is missing or a command argument was passed to the program\n", e)
	fmt.Printf("To change Triage login account or recreate %s remove the %s file, and create a %s file\n", e, e, s)
	fmt.Printf("The format of %s is as follows. Cut and paste then replace account and password.\nThis json file will be deleted on restart of application\n", s)
	fmt.Println(`
{
		"Administrator": {
				"Account":"admin",
				"Password":"abc123"
		}
}`)
	os.Exit(5)

}

// checkJSecureFiles function checks for secure file, and creates it from json file if it doesn't exist.
// both secure file and json are missing, exit.
func checkJSecureFiles(e, s string) {
	if _, err := os.Stat(s); err == nil {
		data, err := ioutil.ReadFile(s)
		encryptFile(e, data, "111222333")
		if err != nil {
			fmt.Print(err)
		}

		err = os.Remove(s)
		if err != nil {
			fmt.Println(err)
		}

	}
	if _, err := os.Stat(e); err != nil {
		Help(e, s)
	}
}

// loadJSecureAccounts function loads account information from secure file
func loadJSecureAccounts(s string) Authorization {
	decrypted := decryptFile(s, "111222333")
	var p Authorization

	err := json.Unmarshal([]byte(decrypted), &p)
	if err != nil {
		fmt.Println("Error!") // error message
	}

	return p
}

func main() {

	// Checking for arguments, and login information
	encfile := ".\\account.enc"
	startup_json_file := ".\\account.json"
	checkJSecureFiles(encfile, startup_json_file)

	// if user passes anything to this program spout out help
	programName := os.Args[0]
	fmt.Println(programName)
	if len(os.Args[1:]) < 1 {
		fmt.Printf("Starting %s\n", programName)
	} else {
		Help(encfile, startup_json_file)
	}

	// Check and load secure accounts
	p := loadJSecureAccounts(encfile)

	fmt.Printf("%s = %s\n", p.Administrator.Account, p.Administrator.Password)

	//ciphertext := encrypt([]byte("Hello World"), "112233")
	//fmt.Printf("Encrypted: %x\n", ciphertext)
	//plaintext := decrypt(ciphertext, "password")
	//fmt.Printf("Decrypted: %s\n", plaintext)

}
