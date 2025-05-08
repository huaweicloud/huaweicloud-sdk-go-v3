package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKafkaMessageRequestBody kafka删除消息请求体
type DeleteKafkaMessageRequestBody struct {

	// 分区消费位点详情
	Partitions *[]PartitionOffsetEntity `json:"partitions,omitempty"`
}

func (o DeleteKafkaMessageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKafkaMessageRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteKafkaMessageRequestBody", string(data)}, " ")
}
