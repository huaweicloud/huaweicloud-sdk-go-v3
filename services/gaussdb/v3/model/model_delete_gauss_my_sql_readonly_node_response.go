package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlReadonlyNodeResponse Response Object
type DeleteGaussMySqlReadonlyNodeResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 订单号，仅包年/包月返回该参数。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGaussMySqlReadonlyNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlReadonlyNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlReadonlyNodeResponse", string(data)}, " ")
}
