package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddReadonlyNodeResponse Response Object
type AddReadonlyNodeResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，仅扩容包年包月实例的节点数量时返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddReadonlyNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddReadonlyNodeResponse struct{}"
	}

	return strings.Join([]string{"AddReadonlyNodeResponse", string(data)}, " ")
}
