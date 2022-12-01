package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSkillOrderInfoResponse struct {

	// 订单总数
	Total *int32 `json:"total,omitempty"`

	// 技能数据
	Data           *[]SkillOrderInfo `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowSkillOrderInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillOrderInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSkillOrderInfoResponse", string(data)}, " ")
}
