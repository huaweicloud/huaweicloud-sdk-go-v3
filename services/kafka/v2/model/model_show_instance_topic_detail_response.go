package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceTopicDetailResponse struct {

	// topic名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 分区列表。
	Partitions *[]ShowInstanceTopicDetailRespPartitions `json:"partitions,omitempty" xml:"partitions"`

	// 订阅该topic的消费组名称列表。
	GroupSubscribed *[]string `json:"group_subscribed,omitempty" xml:"group_subscribed"`
	HttpStatusCode  int       `json:"-"`
}

func (o ShowInstanceTopicDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopicDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopicDetailResponse", string(data)}, " ")
}
