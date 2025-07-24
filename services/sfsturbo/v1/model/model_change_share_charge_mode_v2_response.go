package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeShareChargeModeV2Response Response Object
type ChangeShareChargeModeV2Response struct {

	// 转包周期生成的订单数组
	OrderIds *[]string `json:"order_ids,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeShareChargeModeV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareChargeModeV2Response struct{}"
	}

	return strings.Join([]string{"ChangeShareChargeModeV2Response", string(data)}, " ")
}
