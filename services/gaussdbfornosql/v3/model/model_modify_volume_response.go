package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ModifyVolumeResponse struct {

	// 任务ID，仅按需实例时会返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，仅变更包年包月实例存储容量时返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVolumeResponse struct{}"
	}

	return strings.Join([]string{"ModifyVolumeResponse", string(data)}, " ")
}
