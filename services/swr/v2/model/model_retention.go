package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 镜像老化规则
type Retention struct {
	// 回收规则匹配策略，or

	Algorithm string `json:"algorithm"`
	// ID

	Id int32 `json:"id"`
	// 镜像老化规则

	Rules []Rule `json:"rules"`
	// 保留字段

	Scope string `json:"scope"`
}

func (o Retention) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Retention struct{}"
	}

	return strings.Join([]string{"Retention", string(data)}, " ")
}
