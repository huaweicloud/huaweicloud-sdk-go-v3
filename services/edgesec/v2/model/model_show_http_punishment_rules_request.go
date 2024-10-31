package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpPunishmentRulesRequest Request Object
type ShowHttpPunishmentRulesRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 分页查询参数，第page页
	Page *int32 `json:"page,omitempty"`

	// 分页查询参数，每页显示pagesize条记录
	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ShowHttpPunishmentRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpPunishmentRulesRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpPunishmentRulesRequest", string(data)}, " ")
}
