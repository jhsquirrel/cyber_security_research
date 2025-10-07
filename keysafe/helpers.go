package main

import (
    "fmt"
    "errors"
    "os"
    "encoding/json"
    "bufio"
    "strings"
    "log"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "io"
)

type keyuser struct {
  User string
  Pass string
  Site string
}

func check(e error) {
    if e != nil {
	log.Fatal(e)
    }
}

func getYNUserInput() (int, error, string) {
    var yn, YN string
    n, err := fmt.Scanln(&yn)
    if err != nil {
	return n, err, yn
    }
    YN = strings.ToUpper(yn)
    if YN != "Y" && YN != "N" {
        err = errors.New("Please answer Y/N")
        return n, err, YN
    }
    return n, err, YN
}

func getHiddenUserInput() (int, error, string) {
    var keysafepass string
    // https://stackoverflow.com/questions/30363790/silence-user-input-in-scan-function
    fmt.Println("\033[8m") // Hide input
    n, err := fmt.Scanln(&keysafepass)
    fmt.Println("\033[28m") // Show input
    // check that keysafepass is a sensible length
    if len(keysafepass) < 6 {
        err = errors.New("pass too short (must be > 6 characters") 
    } else if len(keysafepass) > 255 {
        err = errors.New("pass too long (must be < 256 characters")
    }
    return n, err, keysafepass
}

func encryptData(keystorepass string, data string) (string, string) {
    // make hash from keystorepass
    h := sha256.New()
    h.Write([]byte(keystorepass))
    key := h.Sum(nil)
    var plaintext = []byte(data)
    block, err := aes.NewCipher(key)
    if err != nil {
        log.Fatal(err.Error())
    }
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        log.Fatal(err.Error())
    }

    // Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
    nonce := make([]byte, aesgcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	log.Fatal(err.Error())
    }

    ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
    var hexcipher = hex.EncodeToString(ciphertext)
    var hexnonce = hex.EncodeToString(nonce)
    return hexnonce, hexcipher
}

func decryptData(keystorepass string, hexnonce string, hexdata string) string {
    // make hash from keystorepass
    h := sha256.New()
    h.Write([]byte(keystorepass))
    key := h.Sum(nil)

    ciphertext, _ := hex.DecodeString(hexdata)
    nonce, _ := hex.DecodeString(hexnonce)

    block, err := aes.NewCipher(key)
    if err != nil {
        log.Fatal(err.Error())
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
	log.Fatal(err.Error())
    }

    plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        log.Fatal(err.Error())
    }

    return string(plaintext)
}

func createKeyStore(storeName string) {
    f, err := os.Create(storeName)
    check(err)
    defer f.Close()
}

func getDataStore(storeName string, keystorepass string) []keyuser {
    dat, err := os.ReadFile(storeName)
    // read the nonce
    hexnonce := make([]byte, 24)  //aesgcm.NonceSize())
    _ = copy(hexnonce, dat[0:24])
    // read the rest of the data
    hexdata := make([]byte, len(dat) - 24)  //aes.NonceSize())
    _ = copy(hexdata, dat[24:])

    js := decryptData(keystorepass, string(hexnonce), string(hexdata))
    var newstore []keyuser
    err = json.Unmarshal([]byte(js), &newstore)
    check(err)
    return newstore
}

func setDataStore(storeName string, storeData []keyuser, keystorepass string) {
    j, err := json.Marshal(storeData)
    if err != nil {
	log.Fatal(err)
    }
    var js = string(j)
    hexnonce, hexciphertext := encryptData(keystorepass, js)
    // ensure that the file is truncated so no data from previous saves exists
    f, err := os.OpenFile(storeName, os.O_RDWR|os.O_TRUNC, 644)
    defer f.Close()
    check(err)
    w := bufio.NewWriter(f)
    // write the nonce
    _, err = w.WriteString(hexnonce)
    // write the encrypted data
    _, err = w.WriteString(hexciphertext)
    check(err)
    w.Flush()
}
