package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateIssuesRequestBody struct {

	// 分支名
	Branch string `json:"branch"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 工作项集合
	RelatedId []string `json:"related_id"`

	// 仓库id
	RepoId string `json:"repo_id"`
}

func (o AssociateIssuesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIssuesRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateIssuesRequestBody", string(data)}, " ")
}
