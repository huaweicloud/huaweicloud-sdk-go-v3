package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSasTokenResponse Response Object
type CreateSasTokenResponse struct {

	// API key的client_id
	ClientId *string `json:"client_id,omitempty"`

	// 超期时间，UTC格式
	Expiry *string `json:"expiry,omitempty"`

	// 签名字符串
	Signature      *string `json:"signature,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSasTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSasTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateSasTokenResponse", string(data)}, " ")
}
