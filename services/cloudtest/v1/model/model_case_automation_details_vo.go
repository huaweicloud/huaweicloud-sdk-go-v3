package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaseAutomationDetailsVo 用例自动化率详情
type CaseAutomationDetailsVo struct {

	// 用例自动化率
	CaseAutomationRate *string `json:"case_automation_rate,omitempty"`

	// 服务类型对应的用例数目
	ServiceTypeNumberList *[]NameAndValueVo `json:"service_type_number_list,omitempty"`
}

func (o CaseAutomationDetailsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseAutomationDetailsVo struct{}"
	}

	return strings.Join([]string{"CaseAutomationDetailsVo", string(data)}, " ")
}
