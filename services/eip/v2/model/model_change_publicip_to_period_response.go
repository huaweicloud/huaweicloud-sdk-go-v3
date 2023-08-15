package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePublicipToPeriodResponse Response Object
type ChangePublicipToPeriodResponse struct {

	// 转包IP列表
	PublicipIds *[]string `json:"publicip_ids,omitempty"`

	// 订单ID
	OrderId *string `json:"order_id,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangePublicipToPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePublicipToPeriodResponse struct{}"
	}

	return strings.Join([]string{"ChangePublicipToPeriodResponse", string(data)}, " ")
}
