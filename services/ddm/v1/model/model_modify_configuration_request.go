package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyConfigurationRequest Request Object
type ModifyConfigurationRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	Body *ParaGroupUpdate `json:"body,omitempty"`
}

func (o ModifyConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ModifyConfigurationRequest", string(data)}, " ")
}
