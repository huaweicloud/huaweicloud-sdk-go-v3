package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretRequest Request Object
type DeleteSecretRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 密钥ID
	SecretId string `json:"secret_id"`
}

func (o DeleteSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretRequest", string(data)}, " ")
}
