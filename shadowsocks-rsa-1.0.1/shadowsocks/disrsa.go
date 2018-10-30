package shadowsocks

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	//"encoding/base64"
	"encoding/pem"
	"errors"
//	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"log"
)
type MsgConf struct {   	
	Msgpath		string 
	Breakpath	string
}

var decrypted string
var privateKey, publicKey []byte

func BreakMsg(path string)( []byte){// -----------------------------------------------BREAK
	log.Println("------------------------breaking------------------------------------")
	var err error
	privateKey,err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
//		var data []byte
		
		var kikk []byte
	    kikk,err = ioutil.ReadFile(path)
		log.Println("---------------------ok---------------------------------------")

		origData, err:= RsaDecrypt(kikk)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(origData))
		return origData
}

func RsaMsg(path string) {//-----------------------------------------------Encrypt
	var err error
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	
	var data []byte

	var kikk []byte
    kikk,err = ioutil.ReadFile(path)
	data, err = RsaEncrypt(kikk)
	if err != nil {
		panic(err)
	}
	CreateFile(data)
	log.Println(string(data))
}


func CreateFile(data []byte){
	f,err:= os.Create("config.json")
	defer f.Close()
	if err!=nil {
		fmt.Println(err.Error())
	}else{
		_,err:=f.Write(data)
		log.Println(err)
	}
	log.Println(string(data))
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
