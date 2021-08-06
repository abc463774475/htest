package cobratest

import (
	"crypto/rand"
	"fmt"
	"github.com/abc463774475/bbtool/n_log"
	"log"
	"reflect"
	"testing"
)

func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)  // out of randomness, should never happen
	}
	return fmt.Sprintf("%x", buf)
	// or hex.EncodeToString(buf)
	// or base64.StdEncoding.EncodeToString(buf)
}

func TestRand(t *testing.T)  {
	n_log.Info("key is   %v", Key())
}


type ContextOperatorAuth struct {
}

var contextOperatorAuthKey = reflect.TypeOf(ContextOperatorAuth{}).String()

func (req ContextOperatorAuth) ContextKey() string {
	return contextOperatorAuthKey
}


func TestName(t *testing.T) {
	log.Println(contextOperatorAuthKey)
}