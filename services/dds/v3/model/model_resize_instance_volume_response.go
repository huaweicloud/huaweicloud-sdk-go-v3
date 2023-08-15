package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeInstanceVolumeResponse Response Object
type ResizeInstanceVolumeResponse struct {

	// 工作流ID。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，仅扩容包年包月实例的存储容量时返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeInstanceVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceVolumeResponse struct{}"
	}

	return strings.Join([]string{"ResizeInstanceVolumeResponse", string(data)}, " ")
}
