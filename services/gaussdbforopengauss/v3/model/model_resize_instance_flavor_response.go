package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeInstanceFlavorResponse Response Object
type ResizeInstanceFlavorResponse struct {

	// 任务id。按需实例时仅返回任务id。
	JobId *string `json:"job_id,omitempty"`

	// 订单id。仅变更包周期实例会返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeInstanceFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceFlavorResponse struct{}"
	}

	return strings.Join([]string{"ResizeInstanceFlavorResponse", string(data)}, " ")
}
