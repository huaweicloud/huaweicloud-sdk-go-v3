package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResizeColdVolumeResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，仅扩容包年包月实例的存储容量时返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeColdVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeColdVolumeResponse struct{}"
	}

	return strings.Join([]string{"ResizeColdVolumeResponse", string(data)}, " ")
}
