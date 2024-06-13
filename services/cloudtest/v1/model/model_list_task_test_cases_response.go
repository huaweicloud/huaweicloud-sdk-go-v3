package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskTestCasesResponse Response Object
type ListTaskTestCasesResponse struct {

	// 关联的用例uris
	RelatedCaseUris *[]string `json:"related_case_uris,omitempty"`

	// 未关联的用例uris
	NotRelatedCaseUris *[]string `json:"not_related_case_uris,omitempty"`

	// 用例及任务信息
	CaseTaskInfo   *[]RelateTaskTestCasesVo `json:"case_task_info,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListTaskTestCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskTestCasesResponse struct{}"
	}

	return strings.Join([]string{"ListTaskTestCasesResponse", string(data)}, " ")
}
