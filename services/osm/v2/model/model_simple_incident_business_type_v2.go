package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleIncidentBusinessTypeV2 struct {

	// 问题类型id
	BusinessTypeId *string `json:"business_type_id,omitempty"`

	// 问题类型名称
	BusinessTypeName *string `json:"business_type_name,omitempty"`

	// 对应的工单类型0咨询 5报障
	CaseType *string `json:"case_type,omitempty"`

	// 是否可以使用支持计划
	CanUseSupportPlan *bool `json:"can_use_support_plan,omitempty"`
}

func (o SimpleIncidentBusinessTypeV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleIncidentBusinessTypeV2 struct{}"
	}

	return strings.Join([]string{"SimpleIncidentBusinessTypeV2", string(data)}, " ")
}
