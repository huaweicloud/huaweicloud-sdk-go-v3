package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserChargeTypeResponse Response Object
type ShowUserChargeTypeResponse struct {
	Result *ShowUserChargeTypeResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUserChargeTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserChargeTypeResponse struct{}"
	}

	return strings.Join([]string{"ShowUserChargeTypeResponse", string(data)}, " ")
}
