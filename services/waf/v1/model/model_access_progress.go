package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessProgress 冗余参数，仅用于新版console（前端）使用
type AccessProgress struct {

	// 步骤   - 1: 指回源IP加白   - 2: 指本地验证   - 3：指修改DNS解析
	Step *int32 `json:"step,omitempty"`

	// 状态，0：未完成这个步骤；1：已完成这个状态”
	Status *int32 `json:"status,omitempty"`
}

func (o AccessProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessProgress struct{}"
	}

	return strings.Join([]string{"AccessProgress", string(data)}, " ")
}
