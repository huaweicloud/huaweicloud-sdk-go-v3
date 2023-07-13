package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSlowlogDesensitizationRequest Request Object
type SwitchSlowlogDesensitizationRequest struct {

	// 实例ID，可以调用5.3.3 查询实例列表和详情接口获取。如果未申请实例，可以调用5.3.1 创建实例接口创建。
	InstanceId string `json:"instance_id"`

	Body *SwitchSlowlogDesensitizationRequestBody `json:"body,omitempty"`
}

func (o SwitchSlowlogDesensitizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSlowlogDesensitizationRequest struct{}"
	}

	return strings.Join([]string{"SwitchSlowlogDesensitizationRequest", string(data)}, " ")
}
