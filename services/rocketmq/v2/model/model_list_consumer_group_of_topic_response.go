package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConsumerGroupOfTopicResponse Response Object
type ListConsumerGroupOfTopicResponse struct {

	// 消费组列表。
	Groups         *[]string `json:"groups,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListConsumerGroupOfTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupOfTopicResponse struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupOfTopicResponse", string(data)}, " ")
}
