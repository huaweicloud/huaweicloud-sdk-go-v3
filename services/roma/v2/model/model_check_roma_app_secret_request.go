package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRomaAppSecretRequest Request Object
type CheckRomaAppSecretRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o CheckRomaAppSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRomaAppSecretRequest struct{}"
	}

	return strings.Join([]string{"CheckRomaAppSecretRequest", string(data)}, " ")
}
