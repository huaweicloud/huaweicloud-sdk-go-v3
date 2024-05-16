package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSqlLimitRuleOption struct {

	// 最大并发数
	MaxConcurrency int32 `json:"max_concurrency"`

	// 最大等待时间
	MaxWaiting *int32 `json:"max_waiting,omitempty"`
}

func (o UpdateSqlLimitRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitRuleOption struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitRuleOption", string(data)}, " ")
}
