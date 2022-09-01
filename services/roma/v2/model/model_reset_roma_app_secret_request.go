package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetRomaAppSecretRequest struct {

	// 应用ID
	AppId string `json:"app_id" xml:"app_id"`

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *UpdateAppSecretReq `json:"body,omitempty" xml:"body"`
}

func (o ResetRomaAppSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetRomaAppSecretRequest struct{}"
	}

	return strings.Join([]string{"ResetRomaAppSecretRequest", string(data)}, " ")
}
