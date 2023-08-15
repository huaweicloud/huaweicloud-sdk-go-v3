package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSpeedValueResponse Response Object
type SetSpeedValueResponse struct {

	// 业务受理单号
	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SetSpeedValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSpeedValueResponse struct{}"
	}

	return strings.Join([]string{"SetSpeedValueResponse", string(data)}, " ")
}
