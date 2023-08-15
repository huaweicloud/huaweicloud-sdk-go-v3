package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretRequest Request Object
type ShowSecretRequest struct {

	// 密钥ID，从专业版HiLens控制台密钥管理[获取密钥列表](getSecretsListUsingGET.xml)获取
	SecretId string `json:"secret_id"`
}

func (o ShowSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretRequest", string(data)}, " ")
}
