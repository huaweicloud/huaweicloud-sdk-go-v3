package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetSimCardResponse struct {
	// 业务受理单号

	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ResetSimCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetSimCardResponse struct{}"
	}

	return strings.Join([]string{"ResetSimCardResponse", string(data)}, " ")
}
