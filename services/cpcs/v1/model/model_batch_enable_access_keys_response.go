package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableAccessKeysResponse Response Object
type BatchEnableAccessKeysResponse struct {

	// 启用的访问密钥id列表
	AccessKeyIds   *[]string `json:"access_key_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchEnableAccessKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAccessKeysResponse struct{}"
	}

	return strings.Join([]string{"BatchEnableAccessKeysResponse", string(data)}, " ")
}
