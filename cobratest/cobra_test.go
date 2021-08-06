package cobratest

import (
	"github.com/abc463774475/bbtool/n_log"
	_ "github.com/spf13/cobra"

	_ "github.com/gin-gonic/gin"
	"log"
	"testing"
)

func TestCorbra(t *testing.T) {
	log.Println("111")

	n_log.Info("111  %v",11)

	n_log.Erro("222  %v", 22)


	log.SetFlags(log.Lshortfile|log.LstdFlags)
	log.Println("1111")
}
