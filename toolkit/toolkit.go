package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variable of this type will have access
// to all the methods with the reciever *Tools
type Tools struct{}

// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string
func (t *Tools) RandomString(n int) string {
	s := make([]rune, n)            // create a new slice with n length
	r := []rune(randomStringSource) // convert string of bytes into rune

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r)) // random number generate of length r
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
