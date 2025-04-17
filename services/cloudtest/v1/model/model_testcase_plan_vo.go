package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestcasePlanVo 实际的数据类型：单个对象，集合 或 NULL
type TestcasePlanVo struct {

	// 测试用例URI
	TestcaseUri *string `json:"testcase_uri,omitempty"`

	// 测试用例用例编号
	TestcaseNumber *string `json:"testcase_number,omitempty"`

	// 测试计划信息
	Plans *[]TestPlanVo `json:"plans,omitempty"`
}

func (o TestcasePlanVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestcasePlanVo struct{}"
	}

	return strings.Join([]string{"TestcasePlanVo", string(data)}, " ")
}
