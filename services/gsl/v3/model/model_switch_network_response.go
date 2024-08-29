package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchNetworkResponse Response Object
type SwitchNetworkResponse struct {

	// 业务受理单号
	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SwitchNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchNetworkResponse struct{}"
	}

	return strings.Join([]string{"SwitchNetworkResponse", string(data)}, " ")
}
