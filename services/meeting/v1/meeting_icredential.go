package v1

import (
	"bytes"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/auth"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/impl"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/request"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/response"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdkerr"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"time"
)

const (
	AccessTokenInHeader     = "X-Access-Token"
	AccessTokenByAccountUri = "/v1/usg/acs/auth/account"
	AuthorizationInHeader   = "Authorization"
	AccessTokenValidTime    = 3600 * 10
	ClientType              = 72
)

type MeetingCredentials struct {
	UserName      string
	UserPassword  string
	Token         *string
	LastTokenDate *int64
}

type MeetingAuthReqBody struct {
	Account    string `json:"account"`
	ClientType int    `json:"clientType"`
}

type MeetingAuthResBody struct {
	AccessToken string `json:"accessToken"`
}

func (s MeetingCredentials) BuildAccessTokenRequest(endpoint string) *request.DefaultHttpRequest {
	meetingAuthInfo := []byte(s.UserName + ":" + s.UserPassword)
	meetingAuth := "Basic " + base64.StdEncoding.EncodeToString(meetingAuthInfo)
	body := &MeetingAuthReqBody{s.UserName, ClientType}

	return request.NewHttpRequestBuilder().
		WithEndpoint(endpoint).
		WithPath(AccessTokenByAccountUri).
		WithMethod("POST").
		WithBody("MeetingAuthReqBody", body).
		AddHeaderParam(AuthorizationInHeader, meetingAuth).
		AddHeaderParam("Content-Type", "application/json; charset=UTF-8").
		Build()
}

func GetResponseBody(resp *response.DefaultHttpResponse) ([]byte, error) {
	if resp.GetStatusCode() >= 400 {
		return nil, sdkerr.NewServiceResponseError(resp.Response)
	}

	data, err := ioutil.ReadAll(resp.Response.Body)

	if err != nil {
		if closeErr := resp.Response.Body.Close(); closeErr != nil {
			return nil, err
		}
		return nil, err
	}

	if err := resp.Response.Body.Close(); err != nil {
		return nil, err
	} else {
		resp.Response.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	}

	return data, nil
}

func (s MeetingCredentials) ProcessAuthParams(client *impl.DefaultHttpClient, region string) auth.ICredential {
	return s
}

func (s MeetingCredentials) ProcessAuthRequest(client *impl.DefaultHttpClient, req *request.DefaultHttpRequest) (*request.DefaultHttpRequest, error) {
	t := time.Now().Unix()
	if s.Token != nil && *s.Token != "" && t < *s.LastTokenDate+AccessTokenValidTime {
		req.AddHeaderParam(AccessTokenInHeader, *s.Token)
	} else {
		accessTokenReq := s.BuildAccessTokenRequest(req.GetEndpoint())
		accessTokenResp, err := client.SyncInvokeHttp(accessTokenReq)
		if err != nil {
			return nil, err
		}
		data, err := GetResponseBody(accessTokenResp)
		if err != nil {
			return nil, err
		}

		var meetingAuthResInfo MeetingAuthResBody
		err = json.Unmarshal(data, &meetingAuthResInfo)
		if err != nil {
			return nil, err
		}

		s.LastTokenDate = &t
		s.Token = &meetingAuthResInfo.AccessToken

		req.AddHeaderParam(AccessTokenInHeader, *s.Token)
	}

	return req, nil
}

type MeetingCredentialsBuilder struct {
	MeetingCredentials MeetingCredentials
}

func NewMeetingCredentialsBuilder() *MeetingCredentialsBuilder {

	return &MeetingCredentialsBuilder{MeetingCredentials: MeetingCredentials{}}
}

func (builder *MeetingCredentialsBuilder) WithUserName(username string) *MeetingCredentialsBuilder {
	builder.MeetingCredentials.UserName = username
	return builder
}

func (builder *MeetingCredentialsBuilder) WithUserPassword(userPassword string) *MeetingCredentialsBuilder {
	builder.MeetingCredentials.UserPassword = userPassword
	return builder
}

func (builder *MeetingCredentialsBuilder) Build() MeetingCredentials {
	return builder.MeetingCredentials
}
