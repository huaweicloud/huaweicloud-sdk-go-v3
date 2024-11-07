package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceResponse Response Object
type CreateInstanceResponse struct {

	// 实例ID，创建按需付费实例时返回该参数。
	Id *string `json:"id,omitempty"`

	// 订单号，创建包年/包月实例时返回该参数。
	OrderId *string `json:"order_id,omitempty"`

	// 任务id，创建按需付费实例时返回该参数
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceResponse", string(data)}, " ")
}
