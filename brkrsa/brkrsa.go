package main

import (
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "io/ioutil"
    "os"
    "reflect"
    "strconv"
    "math/big"
)

func findFactors(factorsnum *big.Int) []*big.Int{
  zero := big.NewInt(int64(0))
  one := big.NewInt(int64(1))
  i := big.NewInt(int64(1))
  c := big.NewInt(int64(10000000))
  a := big.NewInt(int64(10000000))
  pc := 0
  var primes []*big.Int
  for {
    if i.Cmp(c) > 0 {
      fmt.Println("i=" + i.String() + " factorsnum " + factorsnum.String())
      c = new(big.Int).Add(c, a)
    }
    if i.Cmp(factorsnum) < 0 || i.Cmp(factorsnum) == 0 {
      m := new(big.Int).Mod(factorsnum, i)
      if m.Cmp(zero) == 0 {
        primes = append(primes, i)
        fmt.Println("factor" + i.String())
        pc += 1
      }
      if pc >= 3 {
        // we have found p & q, exit
        return primes
      }
      i = new(big.Int).Add(i, one)
    } else {
      fmt.Println("break")   
      return primes
    }
  }
}

func main() {
    if len(os.Args) != 2 {
        panic(os.Args[0] + " <publickey>")
    }
    pubkeyfile := os.Args[1]
    publicKeyPEM, err := ioutil.ReadFile(pubkeyfile)
    if err != nil {
        panic(err)
    }
    publicKeyBlock, _ := pem.Decode(publicKeyPEM)
    publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
    if err != nil {
        panic(err)
    }
    fmt.Println(reflect.TypeOf(publicKey))
    pk := publicKey.(*rsa.PublicKey)
    N := pk.N
    E := pk.E
    fmt.Println(publicKey)
    fmt.Println("N=" + N.String() + " E=" + strconv.Itoa(E))
    var primes = findFactors(N)
    one := big.NewInt(int64(1))
    var p = new(big.Int).Sub(primes[1], one)
    var q = new(big.Int).Sub(primes[2], one)
    fmt.Println("p=" + primes[1].String())
    fmt.Println("q=" + primes[2].String())
    var pq = new(big.Int).Mul(p, q)
    var e = big.NewInt(int64(E))
    var d = new(big.Int).ModInverse(e, pq)
    fmt.Println("d=" + d.String())
}
