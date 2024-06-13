package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueTreeRequest Request Object
type ListIssueTreeRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 版本id
	VersionId string `json:"version_id"`

	Body *QueryIssueTreeInfo `json:"body,omitempty"`
}

func (o ListIssueTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueTreeRequest struct{}"
	}

	return strings.Join([]string{"ListIssueTreeRequest", string(data)}, " ")
}
