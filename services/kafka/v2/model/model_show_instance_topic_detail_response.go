/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInstanceTopicDetailResponse struct {
	// topic名称。
	Topic *string `json:"topic,omitempty"`
	// 分区列表。
	Partitions *[]ShowInstanceTopicDetailRespPartitions `json:"partitions,omitempty"`
	// 订阅该topic的消费组名称列表。
	GroupSubscribed *[]string `json:"group_subscribed,omitempty"`
	HttpStatusCode  int       `json:"-"`
}

func (o ShowInstanceTopicDetailResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceTopicDetailResponse", string(data)}, " ")
}
