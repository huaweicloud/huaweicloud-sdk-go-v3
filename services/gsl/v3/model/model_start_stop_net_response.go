package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartStopNetResponse struct {

	// 业务受理单号
	WorkOrderId    *int64 `json:"work_order_id,omitempty" xml:"work_order_id"`
	HttpStatusCode int    `json:"-"`
}

func (o StartStopNetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartStopNetResponse struct{}"
	}

	return strings.Join([]string{"StartStopNetResponse", string(data)}, " ")
}
