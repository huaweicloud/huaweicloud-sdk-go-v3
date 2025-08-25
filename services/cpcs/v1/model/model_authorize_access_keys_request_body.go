package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthorizeAccessKeysRequestBody struct {

	// 所需要绑定的应用ID
	AppId string `json:"app_id"`

	// 需要被授权的应用访问密钥的ID列表，若需要授予应用所有访问密钥权限，则仅传入一个元素，且该元素值为“all”，如body为“{\"access_key_ids\": [\"all\"]}”
	AccessKeyIds []string `json:"access_key_ids"`
}

func (o AuthorizeAccessKeysRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeAccessKeysRequestBody struct{}"
	}

	return strings.Join([]string{"AuthorizeAccessKeysRequestBody", string(data)}, " ")
}
