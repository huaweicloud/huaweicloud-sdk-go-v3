package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeAuthorizeAccessKeysRequestBody struct {

	// 需要被解除授权的访问密钥的ID列表，若需要解除所有访问密钥的授权，则仅传入一个元素，且该元素值为“all”，如body为“{\"access_key_ids\": [\"all\"]}”
	AccessKeyIds []string `json:"access_key_ids"`

	// 应用ID
	AppId string `json:"app_id"`
}

func (o DeAuthorizeAccessKeysRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeAuthorizeAccessKeysRequestBody struct{}"
	}

	return strings.Join([]string{"DeAuthorizeAccessKeysRequestBody", string(data)}, " ")
}
