package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisassociateQueueFromEnhancedConnectionRequestBody struct {

	// 需要使用跨源的队列名列表。
	Queues *[]string `json:"queues,omitempty"`

	// 需要使用跨源的弹性资源池名列表。
	ElasticResourcePools *[]string `json:"elastic_resource_pools,omitempty"`
}

func (o DisassociateQueueFromEnhancedConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateQueueFromEnhancedConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociateQueueFromEnhancedConnectionRequestBody", string(data)}, " ")
}
