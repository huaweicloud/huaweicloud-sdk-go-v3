package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAuthorizeAccessKeysResponse Response Object
type CancelAuthorizeAccessKeysResponse struct {

	// 访问密钥ID列表
	AccessKeys     *[]string `json:"access_keys,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CancelAuthorizeAccessKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAuthorizeAccessKeysResponse struct{}"
	}

	return strings.Join([]string{"CancelAuthorizeAccessKeysResponse", string(data)}, " ")
}
