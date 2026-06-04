package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceAliasRequest Request Object
type UpdateInstanceAliasRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateRdsInstanceAliasRequest `json:"body,omitempty"`
}

func (o UpdateInstanceAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAliasRequest", string(data)}, " ")
}
