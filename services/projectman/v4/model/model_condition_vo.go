package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConditionVo struct {

	// 条件类型{&&,||}
	Operator *string `json:"operator,omitempty"`

	// 条件值
	Values *[]string `json:"values,omitempty"`
}

func (o ConditionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionVo struct{}"
	}

	return strings.Join([]string{"ConditionVo", string(data)}, " ")
}
