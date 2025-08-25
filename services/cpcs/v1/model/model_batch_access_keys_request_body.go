package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAccessKeysRequestBody struct {

	// 所需要绑定的访问密钥ID
	AccessKeyIds []string `json:"access_key_ids"`
}

func (o BatchAccessKeysRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAccessKeysRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAccessKeysRequestBody", string(data)}, " ")
}
