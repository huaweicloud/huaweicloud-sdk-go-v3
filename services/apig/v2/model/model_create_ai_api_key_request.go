package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAiApiKeyRequest Request Object
type CreateAiApiKeyRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 凭据编号
	AppId string `json:"app_id"`

	Body *AiApiKeyCreate `json:"body,omitempty"`
}

func (o CreateAiApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiApiKeyRequest struct{}"
	}

	return strings.Join([]string{"CreateAiApiKeyRequest", string(data)}, " ")
}
