package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticForCallDetail struct {

	// 统计对象编号
	Id *string `json:"id,omitempty"`

	// 统计对象名称
	Name *string `json:"name,omitempty"`

	// 调用总量
	CallNum *int32 `json:"call_num,omitempty"`

	// 成功调用量(取数成功)
	SuccessNum *int32 `json:"success_num,omitempty"`

	// 失败调用量(取数失败)
	FailNum *int32 `json:"fail_num,omitempty"`

	// 合法调用量(通过校验)
	LegalNum *int32 `json:"legal_num,omitempty"`

	// 非法调用量(无法通过校验)
	IllegalNum *int32 `json:"illegal_num,omitempty"`
}

func (o StatisticForCallDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticForCallDetail struct{}"
	}

	return strings.Join([]string{"StatisticForCallDetail", string(data)}, " ")
}
