package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodesResponse Response Object
type DeleteNodesResponse struct {

	// 任务ID，仅创建按需实例时会返回该参数。
	JobId *string `json:"job_id,omitempty"`

	// 订单号，创建包年包月时返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodesResponse struct{}"
	}

	return strings.Join([]string{"DeleteNodesResponse", string(data)}, " ")
}
