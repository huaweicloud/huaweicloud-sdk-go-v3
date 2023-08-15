package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpecificationResponse Response Object
type ShowSpecificationResponse struct {

	// 订单列表
	OrderInfos     *[]ProductOrderInfo `json:"orderInfos,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowSpecificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpecificationResponse struct{}"
	}

	return strings.Join([]string{"ShowSpecificationResponse", string(data)}, " ")
}
