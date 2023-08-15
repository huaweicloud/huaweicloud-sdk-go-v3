package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeFlavorResponse Response Object
type ResizeFlavorResponse struct {

	// 规格变更的任务id，仅变更按需实例时会返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 订单id，仅变更包周期实例时会返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeFlavorResponse struct{}"
	}

	return strings.Join([]string{"ResizeFlavorResponse", string(data)}, " ")
}
