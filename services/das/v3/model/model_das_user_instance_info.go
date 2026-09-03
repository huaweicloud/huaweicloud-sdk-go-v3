package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DasUserInstanceInfo DAS实例信息
type DasUserInstanceInfo struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 实例版本号
	EngineVersion *string `json:"engine_version,omitempty"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o DasUserInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DasUserInstanceInfo struct{}"
	}

	return strings.Join([]string{"DasUserInstanceInfo", string(data)}, " ")
}
