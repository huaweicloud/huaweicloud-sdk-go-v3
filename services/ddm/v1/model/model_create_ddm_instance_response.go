package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDdmInstanceResponse Response Object
type CreateDdmInstanceResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 订单号，创建包年/包月实例时返回该参数。
	OrderId *string `json:"order_id,omitempty"`

	// 实例ID。创建按需付费实例时返回该参数。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDdmInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateDdmInstanceResponse", string(data)}, " ")
}
