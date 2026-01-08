package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceGroup 资源分组中的资源信息
type ResourceGroup struct {

	// **参数解释** 资源类型。即命名空间，如弹性云服务器的资源命名空间为：SYS.ECS；各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。 **约束限制** 不涉及 **取值范围** 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。 **默认取值** 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释** 资源的维度信息 **约束限制** 不超过4个维度
	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`

	Status *StatusSchema `json:"status,omitempty"`

	// **参数解释** 事件类型。已废弃。 **取值范围**： 不涉及。
	EventType *int32 `json:"event_type,omitempty"`
}

func (o ResourceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroup struct{}"
	}

	return strings.Join([]string{"ResourceGroup", string(data)}, " ")
}
