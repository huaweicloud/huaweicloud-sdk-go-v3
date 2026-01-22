package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceRecycleInfo struct {

	// **参数解释**： 实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 实例状态。  **约束限制**： 不涉及。 **取值范围**： [详细状态说明请参考[实例状态说明](hrm-api-0010.xml)。](tag:hws,hws_hk,ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,cmcc,hk_tm,dt,cmcc,ax,srg)[详细状态说明请参考[实例状态说明](kafka-api-180514012.xml)。](tag:hcs,fcs,hcs_oemout) **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 实例名称。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消息引擎。  **约束限制**： 不涉及。 **取值范围**： - rocketmq：RocketMQ消息引擎。 - reliability：RocketMQ消息引擎别称。 **默认取值**： 不涉及。
	Engine *string `json:"engine,omitempty"`

	// **参数解释**： 回收时间。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InRecycleTime *int64 `json:"in_recycle_time,omitempty"`

	// **参数解释**： 保留时间。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。 **取值范围**： 1~7。
	SaveTime *int32 `json:"save_time,omitempty"`

	// **参数解释**： 自动删除时间。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AutoDeleteTime *int64 `json:"auto_delete_time,omitempty"`

	// **参数解释**： 每小时的费用。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CostPerHour *float64 `json:"cost_per_hour,omitempty"`

	// **参数解释**： 错误信息。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ErrorMessage *string `json:"error_message,omitempty"`

	// **参数解释**： 产品ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProductId *string `json:"product_id,omitempty"`
}

func (o InstanceRecycleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceRecycleInfo struct{}"
	}

	return strings.Join([]string{"InstanceRecycleInfo", string(data)}, " ")
}
