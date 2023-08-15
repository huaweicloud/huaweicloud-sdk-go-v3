package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSkillOrderInfoRequest Request Object
type ShowSkillOrderInfoRequest struct {

	// 订单ID
	OrderId string `json:"order_id"`
}

func (o ShowSkillOrderInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillOrderInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSkillOrderInfoRequest", string(data)}, " ")
}
