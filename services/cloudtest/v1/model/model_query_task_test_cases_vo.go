package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryTaskTestCasesVo 实际的数据类型：单个对象，集合 或 NULL
type QueryTaskTestCasesVo struct {

	// 关联的用例uris
	RelatedCaseUris *[]string `json:"related_case_uris,omitempty"`

	// 未关联的用例uris
	NotRelatedCaseUris *[]string `json:"not_related_case_uris,omitempty"`

	// 用例及任务信息
	CaseTaskInfo *[]RelateTaskTestCasesVo `json:"case_task_info,omitempty"`
}

func (o QueryTaskTestCasesVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTaskTestCasesVo struct{}"
	}

	return strings.Join([]string{"QueryTaskTestCasesVo", string(data)}, " ")
}
