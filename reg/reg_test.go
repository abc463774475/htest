package reg

import (
	"github.com/abc463774475/bbtool/n_log"
	"github.com/sirupsen/logrus"
	"math/rand"
	"regexp"
	"strconv"
	"testing"
	"time"
)

var (
	reg, err = regexp.Compile("\\$rand\\([0-9]{0,}\\)\\$")
)


func TestReg(t *testing.T)  {
	str := `<d29176df-4a90-4fd7-afc1-a04b393fd220_identifier>=123;Val=$rand(15)$;<3a2a6702-d101-4cd4-9ee4-1ace7b8986cb_identifier>=111`
	// str := "dsx $rand(16)$  ..xd"  v2=$now(2)$<3a2a6702-d101-4cd4-9ee4-1ace7b8986cb_identifier>=222
	for i := 0 ; i < 20; i++{
		go func() {
			str = replaceRand(str)
			n_log.Info("sback  %v", str)
		}()
	}

	time.Sleep(2*time.Second)

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func replaceRand(str string) string {
	sback := reg.ReplaceAllStringFunc(str, func(s string) string {
		startPos := -1
		endPos := -1
		for k := range s {
			if s[k] >= '0' && s[k] <= '9' {
				if startPos == -1 {
					startPos = k
					endPos = k
					continue
				}
				endPos = k
			}
		}
		if startPos == -1 {
			return s
		}

		arg, err := strconv.Atoi(s[startPos:endPos + 1])
		if err != nil {
			logrus.Warnf("s  %v cant not find arg ", s)
			return s
		}

		if arg >= 32 {
			logrus.Warnf("rand argv is so large %v", arg)
			return s
		}

		return RandStringRunes(arg)
	})
	return sback
}

func replaceTime(str string) string {
	reg, err := regexp.Compile("\\$now\\([0-9]{0,}\\)\\$")
	if err != nil {
		logrus.Warn(err)
		return str
	}

	sback := reg.ReplaceAllStringFunc(str, func(s string) string {
		startPos := -1
		endPos := -1
		for k := range s {
			if s[k] >= '0' && s[k] <= '9' {
				if startPos == -1 {
					startPos = k
					endPos = k
					continue
				}
				endPos = k
			}
		}
		if startPos == -1 {
			return s
		}

		arg, err := strconv.Atoi(s[startPos:endPos + 1])
		if err != nil {
			logrus.Warnf("err  %v", err)
		}

		// arg 0 return s   1  return ms 2 return us 3 return ns
		switch arg {
		case 0:
			return strconv.FormatInt(time.Now().Unix(),10)
		case 1:
			return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond),10)
		case 2:
			return strconv.FormatInt(time.Now().UnixNano()/int64(time.Microsecond),10)
		case 3:
			return strconv.FormatInt(time.Now().UnixNano(),10)
		default:
			return s
		}
	})
	return sback
}