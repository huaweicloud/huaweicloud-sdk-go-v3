package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 含阈值告警的QoS数据元素，包括时间，QoS取值，告警状态，产生告警时的阈值。
type QosDataElement struct {

	// QoS时间点, UTC时间，格式：yyyy-MM-ddTHH:mm:ss.SSSZ。
	Time *string `json:"time,omitempty"`

	// QoS值。
	Value *int32 `json:"value,omitempty"`

	// 该时间点是否有阈值告警。 * true: 阈值告警 * false: 无阈值告警
	Alarm *bool `json:"alarm,omitempty"`

	// 该时间点的阈值。
	Threshold *int32 `json:"threshold,omitempty"`
}

func (o QosDataElement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosDataElement struct{}"
	}

	return strings.Join([]string{"QosDataElement", string(data)}, " ")
}
