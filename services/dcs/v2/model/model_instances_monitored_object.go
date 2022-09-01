package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监控对象结构
type InstancesMonitoredObject struct {

	// 测量对象ID，即实例的ID。
	DcsInstanceId *string `json:"dcs_instance_id,omitempty" xml:"dcs_instance_id"`

	// 测量对象名称，即实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 测量对象状态，即实例状态。
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o InstancesMonitoredObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesMonitoredObject struct{}"
	}

	return strings.Join([]string{"InstancesMonitoredObject", string(data)}, " ")
}
