package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfTokenResponse Response Object
type CreateConfTokenResponse struct {
	Data *TokenInfo `json:"data,omitempty"`

	// 企业通讯录查询临时Token。
	AddressToken *string `json:"addressToken,omitempty"`

	// 华为云会议Portal地址。
	GloablPublicIP *string `json:"gloablPublicIP,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConfTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateConfTokenResponse", string(data)}, " ")
}
