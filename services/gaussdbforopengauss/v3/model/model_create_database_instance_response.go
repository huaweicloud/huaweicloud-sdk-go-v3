package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseInstanceResponse Response Object
type CreateDatabaseInstanceResponse struct {
	Instance *CreateDatabaseInstanceResult `json:"instance,omitempty"`

	// 实例创建的任务id。  仅创建按需实例时会返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 创建实例的订单ID。  仅创建包周期实例时会返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDatabaseInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseInstanceResponse", string(data)}, " ")
}
