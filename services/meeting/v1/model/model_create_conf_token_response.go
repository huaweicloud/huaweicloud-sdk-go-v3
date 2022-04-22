package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateConfTokenResponse struct {
	Data *TokenInfo `json:"data,omitempty"`

	// 地址本查询临时Token。
	AddressToken *string `json:"addressToken,omitempty"`

	// global外网IP。
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
