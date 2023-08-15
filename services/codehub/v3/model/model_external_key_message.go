package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExternalKeyMessage struct {

	// 第三方保存在代码托管的关键信息
	ExternalKeyMessage *string `json:"external_key_message,omitempty"`

	// 外部服务名称
	ExternalService *string `json:"external_service,omitempty"`
}

func (o ExternalKeyMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalKeyMessage struct{}"
	}

	return strings.Join([]string{"ExternalKeyMessage", string(data)}, " ")
}
