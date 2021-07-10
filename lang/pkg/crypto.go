package pkg

/*
all the crypto functions are having same interface

New() => this will create a hash.Hash interface
Sum() => this will generate actual hashed value

TODO: so what is the use of New() ??
is New() is used as more or less like salting the hash
*/

import (
	"crypto"
	"crypto/ed25519"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha512"
	"fmt"
)

//hashing
//[learn about hash.Hash type](https://golang.org/pkg/hash/)
//this is an interface which is implemented by some
//New() will give new hash
//Sum() just adds a []byte slice to existing hash (which i will just get from New()) - so now a get => hash + []byte slice

//TODO: so hash function are these - adler32,crc32,crc64,fnv,maphash etc, and md5, sha etc will give one of them when we call md5.new()/sha1.new()
//md5.new()/sha1.new() = any call to say crc32.New() and others
//i don't think so this is correct
//md5.new()/sha1.new() => these give one Way hash
//crc32.New() can be two way (less secure one way we can say)

//md5
//create hashed value from some []byte slice
func CreateMd5Hash() {
	//todo: is it ok just to create hash calling Sum(), benefit of calling New() first
	hash := md5.Sum([]byte("Hello World"))
	fmt.Println(hash)

	//New() generates the hash.Hash - which is an interface
	md := md5.New()
	md5hash := md.Sum([]byte("Hello World"))
	fmt.Println(md5hash)
}

//sha1, sha256, sha512
func CreateShaHash() {
	//sha1
	hash := sha1.New()
	shaHash := hash.Sum([]byte("Hello World"))
	fmt.Println(shaHash)
	//todo: sha256
	//sha512
	hash512 := sha512.New()
	sha512Hash := hash512.Sum([]byte("Hello World"))
	fmt.Println(sha512Hash)
}

//random no generator
//they will be used in some of crypto functions
func RandomNoGenerator() {
	//reader from where we can read random nos - this is an interface which usually gets random no by calling
	//implementations in linux, windows..and other os
	// reader := rand.Reader //rand.Reader => var Reader io.Reader

	//make a slice of []byte in which we will store the generated random no
	sl := make([]byte, 10)
	//read the random nos in the slice
	rand.Read(sl)
	//print slice filled with random no
	fmt.Println(sl)

	//lets take slice of another size
	sl1 := make([]byte, 30)
	rand.Read(sl1)
	fmt.Println(sl1)
}

//aes, des, dsa, ecdsa - just named for knowing i will not work on this as these are less common these days

//rsa - https://www.sohamkamani.com/golang/rsa-encryption/

func RsaFunc() {
	//encryption
	//get rsa publicKey
	randomNoReader := rand.Reader

	//todo: how to convert the stored public key into rsa.PublicKey and use it for encryption
	privateKey, _ := rsa.GenerateKey(randomNoReader, 2048)
	// fmt.Println(privateKey)

	//mesage to be encrypted
	msg := []byte("Hello World")
	publicKey := privateKey.PublicKey
	encryptedData, _ := rsa.EncryptOAEP(sha512.New(), randomNoReader, &publicKey, msg, nil)
	// fmt.Println(encryptedData)

	decryptedBytes, _ := privateKey.Decrypt(nil, encryptedData, &rsa.OAEPOptions{Hash: crypto.SHA512})
	fmt.Println(string(decryptedBytes))

	//todo: using rsa.DecryptOAEP() directly

	//digital signing
	//hash the message before signing so other can not read the message in original
	hash := sha512.New()
	hashData := hash.Sum(msg)

	signature, _ := rsa.SignPSS(randomNoReader, privateKey, crypto.SHA512, hashData, nil)
	verified := rsa.VerifyPSS(&publicKey, crypto.SHA512, hashData, signature, nil)
	fmt.Println(verified)
}

//ed25519
//[why use ed25519 nowdays](https://medium.com/risan/upgrade-your-ssh-key-to-ed25519-c6e8d60d3c54)

func Ed25519Func() {
	//generate a combination of public/private key
	//this function takes a random no Reader - passing nil takes one implemented by go using rand.Reader
	//we will want to pass our own we can pass something else in place of nil
	public, private, _ := ed25519.GenerateKey(nil)

	//sign some message using private key
	message := []byte("Hello World")
	//todo: usually we will hash the message but for this simple example it is ok
	signedBytes := ed25519.Sign(private, message)

	//verify signed bytes using public key - true if signed with corresponding private key
	result := ed25519.Verify(public, message, signedBytes)
	fmt.Println(result) //should be true here
}

//above things in ed25519Func() can be done another way too
func Ed25519Func2() {
	//get the private key
	//seed is 32 bytes long here use any 32 bit []byte slice
	seed := make([]byte, 32)
	//change slice with random no
	rand.Read(seed)
	privateKey := ed25519.NewKeyFromSeed(seed)
	fmt.Println(privateKey)
	//once private key is receive we can get public key or seed used to create the private key
	//get public key
	//todo: here we are getting crypto.PublicKey, how to get ed25519.PrivateKey ??
	//todo: can we get public key frm private key like in rsa to be used in Verify() call
	// publicKey := privateKey.PublicKey

	//now private and public key received call Sign() and Verify() as done above in ed25519Func
	//sign some message using private key
	// message := []byte("Hello World")
	//todo: usually we will hash the message but for this simple example it is ok
	// signedBytes := ed25519.Sign(privateKey, message)

	//verify signed bytes using public key - true if signed with corresponding private key
	// result := ed25519.Verify(publicKey, message, signedBytes)
	// fmt.Println(result) //should be true here
}
