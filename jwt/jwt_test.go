package jwt

import (
	"errors"
	"fmt"
	"github.com/abc463774475/bbtool/n_log"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

var (
	Keys = "774623"
)

type Login struct {
	UID string
	Role int
	StandardClaims jwt.StandardClaims
}

func (l *Login) Valid() error {
	if time.Now().Unix() > l.StandardClaims.ExpiresAt{
		return errors.New("timeooout")
	}
	return nil
}


func GenerateToken(uid string, role int, expireDuration time.Duration) (string,error) {
	expire := time.Now().Add(expireDuration)
	// uid roleid expire regiter
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Login{
		UID:            uid,
		Role:           role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})

	return token.SignedString([]byte(Keys))
}

func TestJwt(t *testing.T)  {
	str,err := GenerateToken("hhxd", 109, 3*time.Second)

	n_log.Info("err  %v  str\n%v",err, str)

	time.Sleep(4*time.Second)
	{
		token,err := jwt.ParseWithClaims(str, &Login{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(Keys), nil
		})
		if err != nil {
			n_log.Erro("err  %v", err)
			return
		}

		l,_ := token.Claims.(*Login)
		n_log.Info("l \n%+v",l)
	}
}

type ClaimsWithClockDrift struct {
	drift int64
	*jwt.StandardClaims
}

func (c *ClaimsWithClockDrift) Valid() error {
	c.StandardClaims.IssuedAt -= c.drift
	valid := c.StandardClaims.Valid()
	c.StandardClaims.IssuedAt += c.drift

	n_log.Info("valid  %v", valid)
	return valid
}

func NewClaimsWithClockDrift(second int64) *ClaimsWithClockDrift {
	return &ClaimsWithClockDrift{drift: second, StandardClaims: &jwt.StandardClaims{}}
}

var (
	publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgMnkWxbrdLCDIILePZ1WZKEf0
Oy5ardL8Qww8hFBtjajQBTCEMxhCiom5WG76uZR9E+w5yuTHaAIWbpip+EyxCSX3
sUJkEXEN6PAeMF06qQJP7BZ56VefEgu4M+BXdRBJiFjT0AvtBxPmGcp6PZ7S15LK
VYsU97D3gdx4DoqhaQIDAQAB
-----END PUBLIC KEY-----`
)

var (
	publicKey1 = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgMnkWxbrdLCDIILePZ1WZKEf0Oy5ardL8Qww8hFBtjajQBTCEMxhCiom5WG76uZR9E+w5yuTHaAIWbpip+EyxCSX3sUJkEXEN6PAeMF06qQJP7BZ56VefEgu4M+BXdRBJiFjT0AvtBxPmGcp6PZ7S15LKVYsU97D3gdx4DoqhaQIDAQAB`
)
func TestJwtDecrypt(t *testing.T)  {
	// token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMzk1Njg3MzIwNjU2NTQ2ODgzIiwiZXhwIjoxNjI1MDQ4NDkyLCJqdGkiOiIxNDEwMTY2NjAzNzk4MDIwMTQ1IiwiaWF0IjoxNjI1MDQ0ODkyLCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJVU0VSIn0.T3p5lRnou0SGsUcz6K3BraAjOLlWjv4DBb4AxK6Snp3fuSPqQx1lwFWx3Xf5zHQ_EQmyIdG75rHQAs9yPathgDP2DyqdQXBUK-CRoJzaHKzEC-LStXx58c-_WFhusecoaC6QB-0oGLDNlq9L6-nmYx0hIXQ-Og65ew6Qq0ND0Qc`
	// token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMzk1Njg3MzIwNjU2NTQ2ODgzIiwiZXhwIjoxNjI1MDUwNzE2LCJqdGkiOiIxNDEwMTc1OTMwNTQxODAzNTY5IiwiaWF0IjoxNjI1MDQ3MTE2LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJVU0VSIn0.mCib78tSIPAas8J7UEMIjDe4m-0i4sIdJB2OcfyDwELAFlRmgCZUE4l5-lji0ectqRCx605Y8DH2UPhn-iexkgEShoe7KN7GIT4_uh6cqWffyr3PKM988ctWJDgOBTBHx9uKRaSXyYb28ZnXxh1PX9252dHD03fO_bkYAiBVgJg`
	token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMzk1Njg3MzIwNjU2NTQ2ODgzIiwiZXhwIjoxNjI1NDU0ODA5LCJqdGkiOiIxNDExODcwODE5NTQxMzIzODMyIiwiaWF0IjoxNjI1NDUxMjA5LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJVU0VSIn0.VUUmbYE32SZc7P5tC7lSk4e80YzAR4JZ6JFQVK3zyUN2X8XipxDC1NiIC9cmc8RmISo1bJEg38qkjJ7_z0BM0SEKohNQjD_SBeylftp61cn-CfliasRtXsuHNV7ZEyXR4vEHP4stApIR_fM2Cwr_UllM7ZVV3-x8_9XygIoyo4Q`

	t1, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		n_log.Erro("err  %v", err)
		return
	}

	claim := t1.Claims.(*ClaimsWithClockDrift).StandardClaims
	n_log.Info("clainm\n%+v", claim)
}