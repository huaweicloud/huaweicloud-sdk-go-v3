package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTriggerResponse struct {

	// 触发动作，update
	Action *string `json:"action,omitempty" xml:"action"`

	// 应用类型，deployments、statefulsets
	AppType *string `json:"app_type,omitempty" xml:"app_type"`

	// 应用名
	Application *string `json:"application,omitempty" xml:"application"`

	// 集群ID（cci时为空）
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id"`

	// 集群名（cci时为空）
	ClusterName *string `json:"cluster_name,omitempty" xml:"cluster_name"`

	// 应用名所在的namespace
	ClusterNs *string `json:"cluster_ns,omitempty" xml:"cluster_ns"`

	// 触发条件，type为all时为.*,type为tag时为tag名,type为regular时为正则表达式
	Condition *string `json:"condition,omitempty" xml:"condition"`

	// 需更新的container名，默认为所有container
	Container *string `json:"container,omitempty" xml:"container"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 创建人
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name"`

	// 是否生效
	Enable *string `json:"enable,omitempty" xml:"enable"`

	// 触发器名
	Name *string `json:"name,omitempty" xml:"name"`

	// 触发器历史
	TriggerHistory *[]TriggerHistories `json:"trigger_history,omitempty" xml:"trigger_history"`

	// 触发器类型，cce、cci
	TriggerMode *string `json:"trigger_mode,omitempty" xml:"trigger_mode"`

	// 触发条件，all、tag、regular
	TriggerType    *string `json:"trigger_type,omitempty" xml:"trigger_type"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTriggerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTriggerResponse struct{}"
	}

	return strings.Join([]string{"ShowTriggerResponse", string(data)}, " ")
}
