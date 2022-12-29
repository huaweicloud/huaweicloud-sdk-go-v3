package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建技能订单请求体
type CreateSkillOrderFrom struct {

	// 技能ID
	SkillId string `json:"skill_id"`

	// 购买个数
	Amount int32 `json:"amount"`

	// 定制技能标识
	CommissionFlag int32 `json:"commission_flag"`
}

func (o CreateSkillOrderFrom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillOrderFrom struct{}"
	}

	return strings.Join([]string{"CreateSkillOrderFrom", string(data)}, " ")
}
