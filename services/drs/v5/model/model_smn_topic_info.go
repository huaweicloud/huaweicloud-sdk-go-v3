package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmnTopicInfo 收件方式与信息体
type SmnTopicInfo struct {

	// Topic的名字。
	Name string `json:"name"`

	// Topic的唯一的资源标识。
	Urn string `json:"urn"`

	// Topic的描述信息。
	Description *string `json:"description,omitempty"`

	// 消息推送的策略，取值： - RETRY_ON_FAILURE：发送失败，保留到失败队列。 - DROP_ON_FAILURE：直接丢弃发送失败的消息。
	PushPolicy string `json:"push_policy"`
}

func (o SmnTopicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnTopicInfo struct{}"
	}

	return strings.Join([]string{"SmnTopicInfo", string(data)}, " ")
}
