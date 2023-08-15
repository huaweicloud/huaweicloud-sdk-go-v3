package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiTestcaseHistoriesRequest Request Object
type ShowApiTestcaseHistoriesRequest struct {

	// 测试用例id
	TestcaseId string `json:"testcase_id"`

	// 起始偏移量，表示从此偏移量开始查询， offset大于等于1
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量,最大支持200条
	Limit *int32 `json:"limit,omitempty"`

	// 测试计划id
	PlanId *string `json:"plan_id,omitempty"`

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ShowApiTestcaseHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiTestcaseHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ShowApiTestcaseHistoriesRequest", string(data)}, " ")
}
