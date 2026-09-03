package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceSimpleDto 实例简单信息
type InstanceSimpleDto struct {

	// 实例ID，实例的唯一标识
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 数据库引擎类型。取值范围：mysql、sqlserver、postgresql、taurus、gaussdbv5、mongodb
	EngineType *string `json:"engine_type,omitempty"`
}

func (o InstanceSimpleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSimpleDto struct{}"
	}

	return strings.Join([]string{"InstanceSimpleDto", string(data)}, " ")
}
