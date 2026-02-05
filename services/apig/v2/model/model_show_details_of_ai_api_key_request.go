package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfAiApiKeyRequest Request Object
type ShowDetailsOfAiApiKeyRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 凭据编号
	AppId string `json:"app_id"`

	// AIAPIKey编号。
	AiApiKeyId string `json:"ai_api_key_id"`
}

func (o ShowDetailsOfAiApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAiApiKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAiApiKeyRequest", string(data)}, " ")
}
