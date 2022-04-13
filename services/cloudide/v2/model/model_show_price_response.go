package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPriceResponse struct {
	// 技术栈价格列表

	Prices *[]ResourcePrice `json:"prices,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPriceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPriceResponse struct{}"
	}

	return strings.Join([]string{"ShowPriceResponse", string(data)}, " ")
}
