package server

import (
	"crypto/sha1"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func (s *Server) sharedKey() ([]byte, error) {
	var b []byte

	switch {
	case s.Params.APIKey != nil:
		b = s.Params.APIKey
	case s.Params.Passphrase != "":
		hash := sha1.Sum([]byte(s.Params.Passphrase))
		b = hash[0:16]
	default:
		b = make([]byte, 16)
		for i := range b {
			b[i] = byte(rand.Intn(256))
		}
	}

	return b, nil
}

func (s *Server) checkAPIKey(sx string) bool {
	k, err := s.sharedKey()
	if err != nil {
		return false
	}
	if sx == fmt.Sprintf("%x", k) {
		return true
	}
	log.Printf("Incorrect api key %q, expected %x", s, k)
	return false
}