package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetSwitchoverRatioInfo struct {

	// **参数解释：** 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 容灾切换的故障节点比例。 **约束限制**： 下限是50，步长是10，最大是100。 **取值范围：** - 50 - 60 - 70 - 80 - 90 - 100 **默认取值：** 100。
	SwitchoverRatio *int32 `json:"switchover_ratio,omitempty"`

	// **参数解释：** 容灾实例数据同步时延，单位s。备实例和主实例同步时延超过该值时，不进行容灾倒换。默认不判断时延。 **约束限制：** 若需指定此参数，最小为10s。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SyncDelay *int64 `json:"sync_delay,omitempty"`
}

func (o SetSwitchoverRatioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSwitchoverRatioInfo struct{}"
	}

	return strings.Join([]string{"SetSwitchoverRatioInfo", string(data)}, " ")
}
