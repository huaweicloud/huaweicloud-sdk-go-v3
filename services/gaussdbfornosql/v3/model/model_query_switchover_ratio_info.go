package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuerySwitchoverRatioInfo struct {

	// **参数解释：** 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 **取值范围：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 容灾切换的故障节点比例。 **取值范围：** - 50 - 60 - 70 - 80 - 90 - 100
	SwitchoverRatio *int32 `json:"switchover_ratio,omitempty"`

	// **参数解释：** 容灾实例数据同步时延，单位s。备实例和主实例同步时延超过该值时，不进行容灾倒换。默认不判断时延。 **取值范围：** 不涉及。
	SyncDelay *int64 `json:"sync_delay,omitempty"`
}

func (o QuerySwitchoverRatioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySwitchoverRatioInfo struct{}"
	}

	return strings.Join([]string{"QuerySwitchoverRatioInfo", string(data)}, " ")
}
