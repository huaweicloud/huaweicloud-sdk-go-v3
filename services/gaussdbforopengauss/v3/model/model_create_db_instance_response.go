package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbInstanceResponse Response Object
type CreateDbInstanceResponse struct {
	Instance *OpenGaussInstanceResult `json:"instance,omitempty"`

	// 实例创建的任务id。  仅创建按需实例时会返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 创建实例的订单ID。  仅创建包周期实例时会返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDbInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateDbInstanceResponse", string(data)}, " ")
}
