package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestcasePlanQueryParam struct {

	// 测试用例URI列表（测试用例URI和编号填其中一个即可）
	CaseUriList *[]string `json:"case_uri_list,omitempty"`

	// 测试用例编号列表（测试用例URI和编号填其中一个即可
	CaseNumberList *[]string `json:"case_number_list,omitempty"`
}

func (o TestcasePlanQueryParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestcasePlanQueryParam struct{}"
	}

	return strings.Join([]string{"TestcasePlanQueryParam", string(data)}, " ")
}
