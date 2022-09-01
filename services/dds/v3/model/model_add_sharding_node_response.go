package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddShardingNodeResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 订单ID，仅扩容包年包月实例的节点数量时返回该参数。
	OrderId        *string `json:"order_id,omitempty" xml:"order_id"`
	HttpStatusCode int     `json:"-"`
}

func (o AddShardingNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddShardingNodeResponse struct{}"
	}

	return strings.Join([]string{"AddShardingNodeResponse", string(data)}, " ")
}
