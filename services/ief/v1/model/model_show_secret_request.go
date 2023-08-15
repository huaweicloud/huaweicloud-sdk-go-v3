package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretRequest Request Object
type ShowSecretRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 密钥ID
	SecretId string `json:"secret_id"`
}

func (o ShowSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretRequest", string(data)}, " ")
}
