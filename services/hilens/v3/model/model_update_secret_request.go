package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretRequest Request Object
type UpdateSecretRequest struct {

	// 密钥ID，从专业版HiLens控制台密钥管理[获取密钥列表](getSecretsListUsingGET.xml)获取
	SecretId string `json:"secret_id"`

	Body *SecretRequestBody `json:"body,omitempty"`
}

func (o UpdateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretRequest", string(data)}, " ")
}
