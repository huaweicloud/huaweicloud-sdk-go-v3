package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveCredentialRequest Request Object
type SaveCredentialRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SaveCredentialRequestBody `json:"body,omitempty"`
}

func (o SaveCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveCredentialRequest struct{}"
	}

	return strings.Join([]string{"SaveCredentialRequest", string(data)}, " ")
}
