package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveOrderResponse Response Object
type SaveOrderResponse struct {

	// 标识
	Id *string `json:"id,omitempty"`

	// 状态，successful/error
	RetStatus      *string `json:"ret_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveOrderResponse struct{}"
	}

	return strings.Join([]string{"SaveOrderResponse", string(data)}, " ")
}
