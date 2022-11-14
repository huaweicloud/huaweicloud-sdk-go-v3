package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchSlowlogDesensitizationRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	// 开启或关闭慢日志脱敏，取值为on或off。
	Status string `json:"status"`
}

func (o SwitchSlowlogDesensitizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSlowlogDesensitizationRequest struct{}"
	}

	return strings.Join([]string{"SwitchSlowlogDesensitizationRequest", string(data)}, " ")
}
