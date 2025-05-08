package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKafkaMessageResponse Response Object
type DeleteKafkaMessageResponse struct {

	// 分区响应信息
	Partitions     *[]PartitionResp `json:"partitions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DeleteKafkaMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKafkaMessageResponse struct{}"
	}

	return strings.Join([]string{"DeleteKafkaMessageResponse", string(data)}, " ")
}
