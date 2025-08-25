package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeAccessKeysResponse Response Object
type AuthorizeAccessKeysResponse struct {

	// 访问密钥ID列表
	AccessKeys     *[]string `json:"access_keys,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AuthorizeAccessKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeAccessKeysResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeAccessKeysResponse", string(data)}, " ")
}
