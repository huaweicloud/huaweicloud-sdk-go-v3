package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaseRealtionInfo struct {

	// 工单id
	CaseId *string `json:"case_id,omitempty" xml:"case_id"`

	// 简要描述
	SimpleDescription *string `json:"simple_description,omitempty" xml:"simple_description"`

	// 提交人，即用户名称
	UserName *string `json:"user_name,omitempty" xml:"user_name"`
}

func (o CaseRealtionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseRealtionInfo struct{}"
	}

	return strings.Join([]string{"CaseRealtionInfo", string(data)}, " ")
}
