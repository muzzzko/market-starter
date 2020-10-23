package password

import (
	"crypto"
	"fmt"
)

type SHA256 struct {
	//todo: add salt

}

func NewSHA256() *SHA256 {
	return &SHA256{}
}

func (s *SHA256) GetHash(text string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(text))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *SHA256) HashIsEqualToPassword(hash string, password string) bool {
	return s.GetHash(password) == hash
}
