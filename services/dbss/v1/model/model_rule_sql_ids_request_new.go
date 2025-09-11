package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleSqlIdsRequestNew struct {

	// 规则ID列表
	Ids []string `json:"ids"`
}

func (o RuleSqlIdsRequestNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleSqlIdsRequestNew struct{}"
	}

	return strings.Join([]string{"RuleSqlIdsRequestNew", string(data)}, " ")
}
