package main

import "fmt"

// comportamiento

// create a struct with the parameters and an interface
type PasswordProtector struct {
	user           string
	passwordName   string
	hashgAlgorithm HashAlgorithm
}

// create an interface with the method(s) to implement
type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

// create a constructor
func NewPasswordProtector(user string, passwordName string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:           user,
		passwordName:   passwordName,
		hashgAlgorithm: hash,
	}
}

// create a function to set the object i want to exchange
func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashgAlgorithm = hash
}

// create a method
func (p *PasswordProtector) Hash() {
	p.hashgAlgorithm.Hash(p)
}

// first intercambiable option
type SHA struct{}

// implementing logic for first option
func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s  \n", p.passwordName)
}

// second intercambiable option
type MD5 struct{}

// implementing logic for second option
func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s  \n", p.passwordName)
}

func main() {
	// create instance for each option
	sha := &SHA{}
	md5 := &MD5{}
	// new single instance
	passwordProtector := NewPasswordProtector("Roosvell", "gmail features", sha)
	passwordProtector.Hash()
	// change environment based on the type of element to handle
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
