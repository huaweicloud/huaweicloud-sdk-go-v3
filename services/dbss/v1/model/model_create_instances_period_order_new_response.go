package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstancesPeriodOrderNewResponse Response Object
type CreateInstancesPeriodOrderNewResponse struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 返回码
	Code *string `json:"code,omitempty"`

	// 订单ID
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstancesPeriodOrderNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancesPeriodOrderNewResponse struct{}"
	}

	return strings.Join([]string{"CreateInstancesPeriodOrderNewResponse", string(data)}, " ")
}
