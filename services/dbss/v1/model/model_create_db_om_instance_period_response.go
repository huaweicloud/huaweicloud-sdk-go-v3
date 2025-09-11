package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbOmInstancePeriodResponse Response Object
type CreateDbOmInstancePeriodResponse struct {

	// 返回码
	Code *string `json:"code,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 订单ID
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDbOmInstancePeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbOmInstancePeriodResponse struct{}"
	}

	return strings.Join([]string{"CreateDbOmInstancePeriodResponse", string(data)}, " ")
}
