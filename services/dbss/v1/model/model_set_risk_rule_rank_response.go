package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRiskRuleRankResponse Response Object
type SetRiskRuleRankResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetRiskRuleRankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRiskRuleRankResponse struct{}"
	}

	return strings.Join([]string{"SetRiskRuleRankResponse", string(data)}, " ")
}
