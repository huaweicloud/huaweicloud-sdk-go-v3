package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RankRiskRuleNew struct {

	// 规则ID
	Id string `json:"id"`

	// 优先级序号，从1开始越小优先级越高
	NewRank int32 `json:"new_rank"`
}

func (o RankRiskRuleNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RankRiskRuleNew struct{}"
	}

	return strings.Join([]string{"RankRiskRuleNew", string(data)}, " ")
}
