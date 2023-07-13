package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretRequest Request Object
type CreateSecretRequest struct {

	// 服务提供者：ief或hilens，选择设备纳管到不同的平台。不填默认为hilens平台
	Provider *string `json:"provider,omitempty"`

	Body *SecretRequestBody `json:"body,omitempty"`
}

func (o CreateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretRequest", string(data)}, " ")
}
