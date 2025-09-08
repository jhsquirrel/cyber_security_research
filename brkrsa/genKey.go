package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "os"
    "strconv"
    "reflect"
    "fmt"
)

func main() {
    if len(os.Args) != 3 {
        panic(os.Args[0] + " <keyname> <keysize>")
    }
    keyname := os.Args[1]
    keysize, err := strconv.Atoi(os.Args[2])
    if err != nil {
        panic(os.Args[1] + " <keysize> not valid")
    }
    pubkeyname := keyname + "_public.pem"
    privkeyname := keyname + "_private.pem"
    privateKey, err := rsa.GenerateKey(rand.Reader, keysize)
    if err != nil {
        panic(err)
    }

    publicKey := &privateKey.PublicKey
    D := privateKey.D
    Primes := privateKey.Primes
    fmt.Println(reflect.TypeOf(privateKey))
    fmt.Println("privateKey D=" + D.String())
    fmt.Println("privateKey Primes=" + Primes[0].String() + " " + Primes[1].String())

    privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
    privateKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: privateKeyBytes,
    })
    err = os.WriteFile(privkeyname, privateKeyPEM, 0644)
    if err != nil {
        panic(err)
    }

    publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        panic(err)
    }
    publicKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: publicKeyBytes,
    })
    err = os.WriteFile(pubkeyname, publicKeyPEM, 0644)
    if err != nil {
        panic(err)
    }
}
