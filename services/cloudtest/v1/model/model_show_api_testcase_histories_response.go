package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiTestcaseHistoriesResponse Response Object
type ShowApiTestcaseHistoriesResponse struct {

	// 测试用例总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 测试服务id
	ProjectId *string `json:"project_id,omitempty"`

	// 测试用例id
	TestcaseId *string `json:"testcase_id,omitempty"`

	// 测试用例名称
	TestcaseName *string `json:"testcase_name,omitempty"`

	// 测试用例结果集
	TestcaseResults *[]TestcaseResult `json:"testcase_results,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ShowApiTestcaseHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiTestcaseHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ShowApiTestcaseHistoriesResponse", string(data)}, " ")
}
