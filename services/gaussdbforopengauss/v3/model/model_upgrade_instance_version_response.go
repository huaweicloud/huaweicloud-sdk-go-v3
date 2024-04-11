package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceVersionResponse Response Object
type UpgradeInstanceVersionResponse struct {

	// 任务id。按需实例时仅返回任务id。
	JobId *string `json:"job_id,omitempty"`

	// 订单id。包周期实例时仅返回订单id。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeInstanceVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceVersionResponse", string(data)}, " ")
}
