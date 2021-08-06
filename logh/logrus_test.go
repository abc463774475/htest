package logh

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"testing"
)

type Output struct {
	f *os.File
}

func (o *Output) Write(p []byte) (n int, err error) {
	o.f.Write(p)
	println(string(p))

	return len(p), nil
}

var (
	GOut = &Output{}
)

func Init()  {
	logrus.SetFormatter(&logrus.JSONFormatter{
	})

	file,err := os.Create("hxd.ini")
	if err != nil {
		log.Fatalf("err %v", err)
	}

	GOut.f = file

	logrus.SetOutput(GOut)
	logrus.SetLevel(logrus.TraceLevel)
}


func TestLogRus(t *testing.T) {
	Init()

	logrus.WithFields(logrus.Fields{
		"name":"xiaodong",
		"age": 18,
	}).Warnf("a good man")
	logrus.Warn("1111111")
	logrus.Errorf("222222")

	logrus.WithField("resType", 10001).Warnf("level warnf")
}

func TestLogrusLevel(t *testing.T) {
	logrus.WithField("resType", 1000).Infof("this is infof")
	logrus.WithField("resType", 10001).Warnf("level warnf")
	logrus.Errorf("this is error")

	logrus.Level()

	logrus.Warn("11111111")
}