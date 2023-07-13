package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretRequest Request Object
type DeleteSecretRequest struct {

	// 密钥ID，从专业版HiLens控制台密钥管理[获密钥列表](getSecretsListUsingGET.xml)获取
	SecretId string `json:"secret_id"`
}

func (o DeleteSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretRequest", string(data)}, " ")
}
