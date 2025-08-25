package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableAccessKeysResponse Response Object
type BatchDisableAccessKeysResponse struct {

	// 禁用的访问密钥id列表
	AccessKeyIds   *[]string `json:"access_key_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDisableAccessKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableAccessKeysResponse struct{}"
	}

	return strings.Join([]string{"BatchDisableAccessKeysResponse", string(data)}, " ")
}
