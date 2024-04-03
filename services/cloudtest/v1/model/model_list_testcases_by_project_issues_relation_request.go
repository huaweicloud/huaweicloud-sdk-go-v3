package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestcasesByProjectIssuesRelationRequest Request Object
type ListTestcasesByProjectIssuesRelationRequest struct {

	// 项目唯一标识，固定长度32位字符
	ProjectId string `json:"project_id"`

	Body *QueryProjectIssuesRelationTestCasesInfo `json:"body,omitempty"`
}

func (o ListTestcasesByProjectIssuesRelationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestcasesByProjectIssuesRelationRequest struct{}"
	}

	return strings.Join([]string{"ListTestcasesByProjectIssuesRelationRequest", string(data)}, " ")
}
