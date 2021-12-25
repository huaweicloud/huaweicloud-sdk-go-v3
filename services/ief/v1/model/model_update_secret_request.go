package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSecretRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
	// 密钥ID

	SecretId string `json:"secret_id"`

	Body *UpdateSecret `json:"body,omitempty"`
}

func (o UpdateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretRequest", string(data)}, " ")
}
