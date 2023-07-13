package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateColdVolumeResponse Response Object
type CreateColdVolumeResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，仅创建包年包月实例的冷数据存储时返回该参数
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateColdVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateColdVolumeResponse struct{}"
	}

	return strings.Join([]string{"CreateColdVolumeResponse", string(data)}, " ")
}
