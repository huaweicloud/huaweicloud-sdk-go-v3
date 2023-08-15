package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSkillOrderListResponse Response Object
type ShowSkillOrderListResponse struct {

	// 订单总数
	Total *int32 `json:"total,omitempty"`

	// 技能数据
	Data           *[]SkillOrderInfo `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowSkillOrderListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillOrderListResponse struct{}"
	}

	return strings.Join([]string{"ShowSkillOrderListResponse", string(data)}, " ")
}
