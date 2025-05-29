package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMongosNodeResponse Response Object
type DeleteMongosNodeResponse struct {

	// 任务ID，仅按需实例返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 订单ID，仅包周期实例返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMongosNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMongosNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteMongosNodeResponse", string(data)}, " ")
}
