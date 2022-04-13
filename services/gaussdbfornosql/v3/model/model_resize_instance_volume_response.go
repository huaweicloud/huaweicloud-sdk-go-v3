package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResizeInstanceVolumeResponse struct {
	// 任务ID，仅按需实例时会返回该参数。

	JobId *string `json:"job_id,omitempty"`
	// 订单ID，仅创建包年包月实例时返回该参数。

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
