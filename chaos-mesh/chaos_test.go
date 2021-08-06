package chaos_mesh

import (
	"fmt"
	"github.com/abc463774475/bbtool/n_log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	// _ "github.com/chaos-mesh/chaos-mesh"
)

func TestChaos(t *testing.T)  {
	strconv.Itoa(111)
}

func f1(s string)  {

}

func BenchmarkChaos(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		s := fmt.Sprint(rand.Int())
		f1(s)
	}
}

func BenchmarkChaos1(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(rand.Int())
		f1(s)
	}
}

func f2([]byte)  {

}

func BenchmarkSlice(b *testing.B)  {
	s1 := "hewllo world"
	for i := 0; i < b.N; i++ {
		s := []byte(s1)
		f2(s)
	}
}

func BenchmarkSlice1(b *testing.B)  {
	s := []byte("hewllo world")
	for i := 0; i < b.N; i++ {
		f2(s)
	}
}

func BenchmarkMap(b *testing.B)  {
	t1, _ := ioutil.ReadFile("./f1.ini")
	names := strings.Split(string(t1),"\n")
	m := make(map[string]int)
	for _,f := range names{
		m[f] =  1
	}
}

func BenchmarkMap1(b *testing.B)  {
	t1, _ := ioutil.ReadFile("./f1.ini")
	names := strings.Split(string(t1),"\n")
	m := make(map[string]int,len(names))
	for _,f := range names{
		m[f] =  1
	}
}

func BenchmarkMapAll(b *testing.B)  {
	b.Run("not make", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t1, _ := ioutil.ReadFile("./f1.ini")
			names := strings.Split(string(t1),"\n")
			m := make(map[string]int)
			for _,f := range names{
				m[f] =  1
			}
		}
	})

	b.Run("make", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t1, _ := ioutil.ReadFile("./f1.ini")
			names := strings.Split(string(t1),"\n")
			m := make(map[string]int,len(names))
			for _,f := range names{
				m[f] =  1
			}
		}
	})
}

func TestWriteFile(b *testing.T)  {
	f1,err := os.OpenFile("f1.ini",os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		n_log.Panic("err  %v", err)
	}
	for i := 0; i < 10000; i++ {
		f1.Write([]byte(fmt.Sprintf("%v\n",i+1)))
	}

	f1.Close()
}

func TestM1(t *testing.T)  {
	return
	tests := []struct{
		give     string
		wantHost string
		wantPort string
	}{
		{
			give:     "192.0.2.0:8000",
			wantHost: "192.0.2.0",
			wantPort: "8000",
		},
		{
			give:     "192.0.2.0:http",
			wantHost: "192.0.2.0",
			wantPort: "http",
		},
		{
			give:     ":8000",
			wantHost: "",
			wantPort: "8000",
		},
		{
			give:     "1:8",
			wantHost: "1",
			wantPort: "8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			host,port, err := net.SplitHostPort(tt.give)
			require.NoError(t, err)

			assert.Equal(t, t, tt.wantHost, host)
			assert.Equal(t, tt.wantPort, port)
		})
	}

}

func f1001() (err error) {
	f,err := ioutil.ReadFile("f2.ini")
	if f == nil {

	}
	return
}

func TestErro(t *testing.T)  {
	err := f1001()
	n_log.Info("err %v", err)
}