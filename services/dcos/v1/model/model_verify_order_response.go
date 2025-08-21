package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyOrderResponse Response Object
type VerifyOrderResponse struct {

	// 标识
	Id *string `json:"id,omitempty"`

	// 状态，successful/error
	RetStatus      *string `json:"ret_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o VerifyOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyOrderResponse struct{}"
	}

	return strings.Join([]string{"VerifyOrderResponse", string(data)}, " ")
}
