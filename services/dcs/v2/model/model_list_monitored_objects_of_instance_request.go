package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMonitoredObjectsOfInstanceRequest struct {

	// 主维度对象ID，与DCS实例列表中实例ID相同。
	InstanceId string `json:"instance_id"`

	// 主维度ID，当前支持子维度的主维度ID的有dcs_instance_id。
	DimName string `json:"dim_name"`
}

func (o ListMonitoredObjectsOfInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitoredObjectsOfInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListMonitoredObjectsOfInstanceRequest", string(data)}, " ")
}
