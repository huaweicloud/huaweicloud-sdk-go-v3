package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShrinkInstanceNodeResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 订单ID，仅创建包年包月实例时返回该参数。
	OrderId        *string `json:"order_id,omitempty" xml:"order_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkInstanceNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodeResponse", string(data)}, " ")
}
