package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrepayResizeResponse Response Object
type CreatePrepayResizeResponse struct {

	// 订单编号
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrepayResizeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrepayResizeResponse struct{}"
	}

	return strings.Join([]string{"CreatePrepayResizeResponse", string(data)}, " ")
}
