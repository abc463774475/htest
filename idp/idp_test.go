package idp
//
//import (
//	"bytes"
//	"context"
//	"fmt"
//	idp_client "git.querycap.com/idp/idp_client/v2"
//	"git.querycap.com/tools/svcutil/confhttp"
//	"github.com/abc463774475/bbtool/n_log"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/go-courier/courier"
//	"github.com/go-courier/httptransport/client"
//	"github.com/go-courier/metax"
//	"net/url"
//	"strings"
//	"testing"
//)
//
//var (
//	IdpMgr    *idp_client.IdpMgr
//	IdpConfig = &idp_client.IdpConfig{
//		ClientEndpoint: confhttp.ClientEndpoint{
//			Client: client.Client{
//				Protocol:              "https",
//				Host:                  "srv-idp---idp.hw-dev.rktl.xyz",
//				Port:                  0,
//			},
//		},
//		Issuer: "idp.rockontrol.com",
//		Mod:    "srv-resource",
//		//Secret: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxNDEwMTg5MDgxOTQ1NTc5NTc2IiwiZXhwIjoxNjMzNzc1Nzk3LCJqdGkiOiIxNDEwNTQ3ODg0ODU5ODUxODMyIiwiaWF0IjoxNjI1MTM1Nzk3LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJDTElFTlQifQ.CCfN_VndTvYHgdVvmvuooaAgsSo_8BDt61LsRAUOXsw-5nWbSnZTy3C7dkOf-yJc2tFpck2-aCMl0tcaZb738mZwSdpefOZAO9Jvfyz5SnAXJA7ndfJVSMvw5_DxsOVfZhCU9QWJ4vQOea6ivayw-n11mk0JfSIDaA5iv_UQFhQ",
//		//Secret: `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxNDEwMTg5MDgxOTQ1NTc5NTc2IiwiZXhwIjoxNjMzODM0MDAzLCJqdGkiOiIxNDEwNzkyMDIxNTIxNjAyNjA5IiwiaWF0IjoxNjI1MTk0MDAzLCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJDTElFTlQifQ.DVLHJx81YR8LucMBg67sy6YEJosImvBGF_zuMnfyaYr8eirgS2wgoSAa_dKj44UCECosPn4kXDRxd91bHtPaAKiVg7aQ9DRkvF2WVm4oAqGjfsd-364iYi5qf75ypfaDyDG9YKlIAsZZeERA9r7rhlZbgEO6CNw4tNVE2dFAsE0`,
//		//Secret: `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxNDEwMTg5MDgxOTQ1NTc5NTc2IiwiZXhwIjoxNjMzODM0MDAzLCJqdGkiOiIxNDEwNzkyMDIxNTIxNjAyNjA5IiwiaWF0IjoxNjI1MTk0MDAzLCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJDTElFTlQifQ.DVLHJx81YR8LucMBg67sy6YEJosImvBGF_zuMnfyaYr8eirgS2wgoSAa_dKj44UCECosPn4kXDRxd91bHtPaAKiVg7aQ9DRkvF2WVm4oAqGjfsd-364iYi5qf75ypfaDyDG9YKlIAsZZeERA9r7rhlZbgEO6CNw4tNVE2dFAsE0`,
//		Secret: `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxNDEwMTg5MDgxOTQ1NTc5NTc2IiwiZXhwIjoxNjMzNzc1Nzk3LCJqdGkiOiIxNDEwNTQ3ODg0ODU5ODUxODMyIiwiaWF0IjoxNjI1MTM1Nzk3LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJDTElFTlQifQ.CCfN_VndTvYHgdVvmvuooaAgsSo_8BDt61LsRAUOXsw-5nWbSnZTy3C7dkOf-yJc2tFpck2-aCMl0tcaZb738mZwSdpefOZAO9Jvfyz5SnAXJA7ndfJVSMvw5_DxsOVfZhCU9QWJ4vQOea6ivayw-n11mk0JfSIDaA5iv_UQFhQ`,
//	}
//
//	//OperationMap = idp_client.InitRoutes(courier.NewRouter(httptransport.Group("/iot-core-proxy-gateway")).Routes())
//)
//
//type IDPClient struct {
//	*idp_client.ClientIdpStruct
//}
//
//func (idp *IDPClient) CurrentUser(req *CurrentUser, metas ...courier.Metadata)  (*AccountCompleteInfo, courier.Metadata, error){
//	return req.InvokeContext(idp.Context(), idp.Client, metas...)
//}
//
//type CurrentUser struct {
//	//Authorization string           `in:"header" name:"Authorization,omitempty"`
//	//AccountIDs    []datatypes.UUID `in:"query" name:"accountID,omitempty"`
//	// Bearer access_token
//	Authorization string `in:"header" name:"Authorization,omitempty"`
//	// Bearer access_token in query
//	AuthorizationInQuery string `in:"query" name:"authorization,omitempty"`
//	Env                  string `in:"query" name:"env,omitempty"`
//}
//
//func (CurrentUser)Path() string {
//	return "/idp/v0/user"
//}
//
//func (CurrentUser) Method() string {
//	return "GET"
//}
//
//type Identities struct {
//	EMAIL  string
//	MOBILE string
//	NICKNAME string
//}
//
//type AccountCompleteInfo struct {
//	idp_client.IdPAccount
//	Identities Identities `json:"identities"`
//}
//
//func (req *CurrentUser) InvokeContext(ctx context.Context, c courier.Client, metas ...courier.Metadata) (*AccountCompleteInfo, courier.Metadata, error) {
//	resp := new(AccountCompleteInfo)
//
//	ctx = metax.ContextWith(ctx, "operationID", "idp.ListAccount")
//	meta, err := c.Do(ctx, req, metas...).Into(resp)
//	return resp, meta, err
//}
//
//func (req *CurrentUser) Invoke(c courier.Client, metas ...courier.Metadata) (*AccountCompleteInfo, courier.Metadata, error) {
//	return req.InvokeContext(context.Background(), c, metas...)
//}
//
//func NewIDPClient(c courier.Client) *IDPClient {
//	ret := &IDPClient{}
//	ret.ClientIdpStruct = idp_client.NewClientIdp(c)
//	return ret
//}
//
//func Init() {
//	n_log.Info("222")
//	IdpConfig.Client.SetDefaults()
//	IdpMgr = idp_client.NewIdpMgr(IdpConfig)
//}
//
//func TestIDPClient(t *testing.T) {
//	n_log.Info("1111")
//	Init()
//	err := IdpMgr.Init(nil)
//	n_log.Info("err  %v", err)
//
//	atoken, err := IdpMgr.GetToken(idp_client.ExchangeTokenParam{
//		Code:         "1394568d5ec00038",
//		GrantType:    "authorization_code",
//		RefreshToken: "",
//	})
//
//	if err != nil {
//		n_log.Erro("err  %v", err)
//		return
//	}
//	n_log.Info("atoken\n%+v", atoken)
//}
//
//func TestIDP(t *testing.T)  {
//	var client = &client.Client{
//		Host:     "srv-idp---idp.hw-dev.rktl.xyz",
//		Protocol: "https",
//	}
//	client.SetDefaults()
//	idpClient := idp_client.NewClientIdp(client)
//
//	aes, _,err := idpClient.GetToken(&idp_client.GetToken{
//		Authorization:        "bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxNDEwMTg5MDgxOTQ1NTc5NTc2IiwiZXhwIjoxNjMzNzc1Nzk3LCJqdGkiOiIxNDEwNTQ3ODg0ODU5ODUxODMyIiwiaWF0IjoxNjI1MTM1Nzk3LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJDTElFTlQifQ.CCfN_VndTvYHgdVvmvuooaAgsSo_8BDt61LsRAUOXsw-5nWbSnZTy3C7dkOf-yJc2tFpck2-aCMl0tcaZb738mZwSdpefOZAO9Jvfyz5SnAXJA7ndfJVSMvw5_DxsOVfZhCU9QWJ4vQOea6ivayw-n11mk0JfSIDaA5iv_UQFhQ",
//		AuthorizationInQuery: "",
//		Data: struct {
//			Code         string `name:"code,omitempty"`
//			GrantType    string `default:"authorization_code" name:"grant_type,omitempty" validate:"@string{authorization_code,refresh_token}"`
//			RefreshToken string `name:"refresh_token,omitempty"`
//		}{"139458c27a801c38", "authorization_code",""},
//	})
//
//	if err != nil {
//		n_log.Panic("err  %v", err)
//		return
//	}
//
//	n_log.Info("aes  \n%+v", aes)
//}
//
//func TestList(t *testing.T)  {
//	var client = &client.Client{
//		Host:     "srv-idp---idp.hw-dev.rktl.xyz",
//		Protocol: "https",
//	}
//	client.SetDefaults()
//	idpClient := NewIDPClient(client)
//
//	at,_, err := idpClient.ListAccount(&idp_client.ListAccount{
//		Authorization: `bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMzk1Njg3MzIwNjU2NTQ2ODgzIiwiZXhwIjoxNjI1MjEzNjg3LCJqdGkiOiIxNDEwODU5NDgxOTgxNTI1MDQ4IiwiaWF0IjoxNjI1MjEwMDg3LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJVU0VSIn0.xoaQ5zM4AM2DpvMjj7DhkMPex8AgdYfRSu74Nwg-GVGRvRhjxoOzabibh1VmRZvwpyFYp-qLjrV-ClbUffNyvpl1Fv_sQrStTD7bU6XjbrfwN_o6uXLev9hk2VQoNL-YjtuTm3hoQCPWtuzA6Ie2jMh2GAo4r0xzhRoLWb4uttc`,
//		AccountIDs:    nil,
//	})
//
//	if err != nil {
//		n_log.Panic("err  %v", err)
//		return
//	}
//
//	n_log.Info("aes  \n%+v", at)
//}
//
//func TestCurrentUser(t *testing.T)  {
//	var client = &client.Client{
//		Host:     "srv-idp---idp.hw-dev.rktl.xyz",
//		Protocol: "https",
//	}
//	client.SetDefaults()
//	idpClient := NewIDPClient(client)
//
//	at,_, err := idpClient.CurrentUser(&CurrentUser{
//		Authorization: `bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMzk1Njg3MzIwNjU2NTQ2ODgzIiwiZXhwIjoxNjI1MjEzNjg3LCJqdGkiOiIxNDEwODU5NDgxOTgxNTI1MDQ4IiwiaWF0IjoxNjI1MjEwMDg3LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJVU0VSIn0.xoaQ5zM4AM2DpvMjj7DhkMPex8AgdYfRSu74Nwg-GVGRvRhjxoOzabibh1VmRZvwpyFYp-qLjrV-ClbUffNyvpl1Fv_sQrStTD7bU6XjbrfwN_o6uXLev9hk2VQoNL-YjtuTm3hoQCPWtuzA6Ie2jMh2GAo4r0xzhRoLWb4uttc`,
//		Env: "online",
//	})
//
//	if err != nil {
//		n_log.Panic("err  %v", err)
//		return
//	}
//
//	n_log.Info("aes  \n%+v", at)
//}
//
//type ClaimsWithClockDrift struct {
//	drift int64
//	*jwt.StandardClaims
//}
//
//func (c *ClaimsWithClockDrift) Valid() error {
//	//c.StandardClaims.IssuedAt -= c.drift
//	//valid := c.StandardClaims.Valid()
//	//c.StandardClaims.IssuedAt += c.drift
//	//
//	//n_log.Info("valid  %v", valid)
//	// return valid
//	return nil
//}
//
//func NewClaimsWithClockDrift(second int64) *ClaimsWithClockDrift {
//	return &ClaimsWithClockDrift{drift: second, StandardClaims: &jwt.StandardClaims{}}
//}
//
//
//func TestPublick(t *testing.T)  {
//	var client = &client.Client{
//		Host:     "srv-idp---idp.hw-dev.rktl.xyz",
//		Protocol: "https",
//	}
//	client.SetDefaults()
//	idpClient := NewIDPClient(client)
//
//	publicKey,_, err := idpClient.RSAPublicKey()
//	if err != nil {
//		n_log.Erro("err  %v", err)
//	}
//
//	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(*publicKey))
//	if err != nil {
//		n_log.Erro("err  %v", err)
//	}
//
//	// token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMzk1Njg3MzIwNjU2NTQ2ODgzIiwiZXhwIjoxNjI1NDU0ODA5LCJqdGkiOiIxNDExODcwODE5NTQxMzIzODMyIiwiaWF0IjoxNjI1NDUxMjA5LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJVU0VSIn0.VUUmbYE32SZc7P5tC7lSk4e80YzAR4JZ6JFQVK3zyUN2X8XipxDC1NiIC9cmc8RmISo1bJEg38qkjJ7_z0BM0SEKohNQjD_SBeylftp61cn-CfliasRtXsuHNV7ZEyXR4vEHP4stApIR_fM2Cwr_UllM7ZVV3-x8_9XygIoyo4Q`
//	token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMzk1Njg3MzIwNjU2NTQ2ODgzIiwiZXhwIjoxNjI1NDU3NTg3LCJqdGkiOiIxNDExODgyNDY5OTg0MDQ4MTg0IiwiaWF0IjoxNjI1NDUzOTg3LCJpc3MiOiJpZHAucm9ja29udHJvbC5jb20iLCJzdWIiOiJVU0VSIn0.U9fXxSWlXRyMU1iZerhVytACk7wjkNUCUbhLqFSP-XBD2Iq0FKXrzZZFtkFeHDQmqo5P7nO6eIhGKbu5_xcSjFvyjXpZKk9LNcbOWiBEx7V3Piaw4Gm6Dl7o96Ka7fTy0PSEbgE9V9zqyYTRoUatnAlseujv3_YcYaXlMyCQSr0`
//
//	t1, err := jwt.ParseWithClaims(token, NewClaimsWithClockDrift(3), func(token *jwt.Token) (interface{}, error) {
//
//		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return key, nil
//	})
//
//	if err != nil {
//		n_log.Erro("err  %v", err)
//		return
//	}
//
//
//	n_log.Info("t1  %+v",t1)
//}
//
//var (
//	idpAuthURL = `https://web-idp---idp.hw-dev.querycap.com/authorize`
//	idpClientID = `1410189081945579576`
//)
//
//func authCodeURL(redirectURI, state string) string {
//	var buf bytes.Buffer
//	buf.WriteString(idpAuthURL)
//	v := url.Values{
//		"response_type": {"code"},
//		"client_id":     {idpClientID},
//	}
//	if redirectURI != "" {
//		v.Set("redirect_uri", redirectURI)
//	}
//	if state != "" {
//		v.Set("state", state)
//	}
//	if strings.Contains(idpAuthURL, "?") {
//		buf.WriteByte('&')
//	} else {
//		buf.WriteByte('?')
//	}
//	buf.WriteString(v.Encode())
//	return buf.String()
//}
//
//
//func TestAuth(t *testing.T)  {
//	str := authCodeURL("https://openapi__infra.rocktl.com/srv-resource---iot.hw-iot-dev.rktl.xyz/resource",
//		"Code",
//	)
//
//	n_log.Info("str\n%v\n",str)
//}
