package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnableSimCardResponse struct {
	// 业务受理单号

	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o EnableSimCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableSimCardResponse struct{}"
	}

	return strings.Join([]string{"EnableSimCardResponse", string(data)}, " ")
}
