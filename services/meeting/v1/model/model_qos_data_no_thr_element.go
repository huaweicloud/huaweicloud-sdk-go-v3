package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QosDataNoThrElement 不含阈值告警的QoS数据元素，包括时间，QoS取值。
type QosDataNoThrElement struct {

	// Qos时间点, UTC时间，格式：yyyy-MM-ddTHH:mm:ss.SSSZ。
	Time *string `json:"time,omitempty"`

	// QoS值。
	Value *string `json:"value,omitempty"`
}

func (o QosDataNoThrElement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosDataNoThrElement struct{}"
	}

	return strings.Join([]string{"QosDataNoThrElement", string(data)}, " ")
}
