package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKafkaTopicMessagesResponse Response Object
type DeleteKafkaTopicMessagesResponse struct {

	// 分区响应信息
	Partitions     *[]PartitionResp `json:"partitions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DeleteKafkaTopicMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKafkaTopicMessagesResponse struct{}"
	}

	return strings.Join([]string{"DeleteKafkaTopicMessagesResponse", string(data)}, " ")
}
