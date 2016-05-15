package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "fmt"
    "io/ioutil"
    )

func main() {
    
    var pub *rsa.PublicKey
    random := rand.Reader
    data, err := ioutil.ReadFile("./public.pem")
    
    if err != nil {
        panic(err)
    }
    
    block, _ := pem.Decode(data)
    msg := []byte("hello")
    
    pubInterface, parseErr := x509.ParsePKIXPublicKey(block.Bytes)
    
    if parseErr != nil {
        fmt.Println("Load public key error")
        panic(parseErr)
    }
    
    pub = pubInterface.(*rsa.PublicKey)
    
    encryptedData, encryptedErr := rsa.EncryptPKCS1v15(random, pub, msg)
    fmt.Println(encryptedErr)
    fmt.Println(base64.URLEncoding.EncodeToString(encryptedData))
}