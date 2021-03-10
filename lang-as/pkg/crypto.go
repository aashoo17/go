package main

/*
all the crypto functions are having same interface

New() => this will create a hash.Hash interface
Sum() => this will generate actual hashed value, so only calling Sum() is sufficient to generate hash

TODO: so what is the use of New() ??
is New() is used as more or less like salting the hash

*/

import (
	"crypto/ed25519"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha512"
	"fmt"
)

func main() {

}

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
func createMd5Hash() {
	hash := md5.Sum([]byte("Hello World"))
	fmt.Println(hash)
}

func createMd5Hash2() {
	//New() generates the hash.Hash - which is an interface
	md := md5.New()
	// fmt.Println(md)
	// Sum appends the current hash to byte passed in Sum([]byte) and returns the resulting slice.
	// It does not change the underlying hash state.
	//[hash.New() reference](https://golang.org/pkg/hash/#Hash)
	hash := md.Sum([]byte("Hello World"))
	fmt.Println(hash)
}

//sha1, sha256, sha512
func createShaHash() {
	hash := sha1.New()
	shaHash := hash.Sum([]byte("Hello World"))
	fmt.Println(shaHash)

	//sha512
	hash512 := sha512.New()
	sha512Hash := hash512.Sum([]byte("Hello World"))
	fmt.Println(sha512Hash)
}

//aes, des, dsa, ecdsa - just named for knowing i will not work on this as these are less common these days

//public - private key cryptography
//in authenticating login etc - identity verification
//here we pretty much get public key & private key
//so generation of both is first criteria
//verifying private key with public key is second and that's all we will want generally

//in encrypting, decrypting data
//create public, private keys
//or this can be used like this
//encrypt some mesage with public key
//decrypt the message using private key

//random no generator
//they will be used in some of crypto functions
func randomNoGenerator() {
	//reader from where we can read random nos - this is an interface which usually gets random no by calling
	//implementaions in linux, windows..and other os
	reader := rand.Reader

	//make a slice of []byte in which we will store the generated random no
	sl := make([]byte, 10)
	//read the random nos in the slice
	reader.Read(sl)
	//print slice filled with random no
	fmt.Println(sl)

	//lets take slice of another size
	sl1 := make([]byte, 30)
	reader.Read(sl1)
	fmt.Println(sl1)
}

//rsa

//OAEP and PSS funcs from go should now used hence i am just using them - read the golang docs to know why
func rsaFunc() {

	//encrption
	//get rsa publicKey
	randomNoReader := rand.Reader
	publicKey := rsa.PublicKey{}
	fmt.Println(publicKey)

	//mesage to be encrypted
	msg := []byte("Hello World")
	label := []byte("order")

	rsa.EncryptOAEP(sha512.New(), randomNoReader, &publicKey, msg, label)
}

//ed25519
//[why use ed25519 nowdays](https://medium.com/risan/upgrade-your-ssh-key-to-ed25519-c6e8d60d3c54)

func ed25519Func() {
	//generate a combination of public/private key
	//this function takes a random no Reader - passing nil takes one implemented by go using rand.Reader
	//we will want to pass our own we can pass something else in place of nil
	public, private, _ := ed25519.GenerateKey(nil)

	//sign some message using private key
	message := []byte("Hello World")
	signedBytes := ed25519.Sign(private, message)

	//verify signed bytes using public key - true if signed with corresponding private key
	result := ed25519.Verify(public, message, signedBytes)
	fmt.Println(result) //should be true here
}

//above things in ed25519Func() can be done another way too
func ed25519Func2() {
	//get the private key
	//seed is 32 bytes long here use any 32 bit []byte slice
	seed := make([]byte, 32)
	//change a bit as default one is having all 0 in byte slice
	seed[0] = 10
	privateKey := ed25519.NewKeyFromSeed(seed)

	//once private key is receive we can get public key or seed used to create the private key
	//get public key
	publicKey := privateKey.Public()
	seedUsedForPrivateKey := privateKey.Seed()

	fmt.Println(publicKey)
	fmt.Println(seedUsedForPrivateKey)

	//now private and public key received call Sign() and Verify() as done above in ed25519Func
}

//TODO: explain from above ed25519Func() and ed25519Func2() what can be done using ed25519 cryto func
//so an someone having - publicKey and privateKey
//i can sign bytes to get new bytes - called as signed bytes using my private key
//now if someone is having publicKey he can check if it was signed by someone with this privateKey
//so if we can get the original bytes which was converted when signed using privateKey - by using publicKey which was not in api
//then encryption has no meaning
