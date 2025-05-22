package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckSecretsRequestBody struct {

	// 凭据检测类型。
	Type string `json:"type"`

	// 凭据检测对应的请求数据体。
	Data *interface{} `json:"data"`
}

func (o CheckSecretsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSecretsRequestBody struct{}"
	}

	return strings.Join([]string{"CheckSecretsRequestBody", string(data)}, " ")
}
