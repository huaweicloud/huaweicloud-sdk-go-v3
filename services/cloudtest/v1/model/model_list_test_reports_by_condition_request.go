package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestReportsByConditionRequest Request Object
type ListTestReportsByConditionRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 每页显示的条目数量,最大支持200条
	PageSize int64 `json:"page_size"`

	// 页数，page_no大于等于1
	Offset int64 `json:"offset"`

	// 名称关键词
	KeyWord *string `json:"key_word,omitempty"`

	// 是否是我的测试报告
	Own *bool `json:"own,omitempty"`
}

func (o ListTestReportsByConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestReportsByConditionRequest struct{}"
	}

	return strings.Join([]string{"ListTestReportsByConditionRequest", string(data)}, " ")
}
