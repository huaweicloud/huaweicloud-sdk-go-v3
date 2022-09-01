package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Trigger struct {

	// 触发动作，update
	Action string `json:"action" xml:"action"`

	// 应用类型，deployments、statefulsets
	AppType string `json:"app_type" xml:"app_type"`

	// 应用名
	Application string `json:"application" xml:"application"`

	// 集群ID（cci时为空）
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// 集群名（cci时为空）
	ClusterName string `json:"cluster_name" xml:"cluster_name"`

	// 应用名所在的namespace
	ClusterNs string `json:"cluster_ns" xml:"cluster_ns"`

	// 触发条件，type为all时为.*,type为tag时为tag名,type为regular时为正则表达式
	Condition string `json:"condition" xml:"condition"`

	// 需更新的container名，默认为所有container
	Container string `json:"container" xml:"container"`

	// 创建时间
	CreatedAt string `json:"created_at" xml:"created_at"`

	// 创建人
	CreatorName string `json:"creator_name" xml:"creator_name"`

	// 是否生效
	Enable string `json:"enable" xml:"enable"`

	// 触发器名
	Name string `json:"name" xml:"name"`

	// 触发器历史
	TriggerHistory []TriggerHistories `json:"trigger_history" xml:"trigger_history"`

	// 触发器类型，cce、cci
	TriggerMode string `json:"trigger_mode" xml:"trigger_mode"`

	// 触发条件，all、tag、regular
	TriggerType string `json:"trigger_type" xml:"trigger_type"`
}

func (o Trigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}
