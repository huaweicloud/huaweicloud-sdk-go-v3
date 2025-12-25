package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataPointInfo struct {

	// **参数解释**： 计算出该条告警记录的资源监控数据上报的UTC时间。 **取值范围**： 字符串长度在 1 到 64 之间。
	Time *string `json:"time,omitempty"`

	// **参数解释**： 计算出该条告警记录的资源监控数据在该时间点的监控数值，如：7.019 **取值范围**： 整数，最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108
	Value *float64 `json:"value,omitempty"`
}

func (o DataPointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPointInfo struct{}"
	}

	return strings.Join([]string{"DataPointInfo", string(data)}, " ")
}
