package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetailTestTypeCasePassRateVo 每个测试类型的用例通过率
type DetailTestTypeCasePassRateVo struct {

	// 测试类型
	TestType *int32 `json:"test_type,omitempty"`

	// 用例通过率
	CasePassRate *string `json:"case_pass_rate,omitempty"`
}

func (o DetailTestTypeCasePassRateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailTestTypeCasePassRateVo struct{}"
	}

	return strings.Join([]string{"DetailTestTypeCasePassRateVo", string(data)}, " ")
}
