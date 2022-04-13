package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCommitsRequest struct {
	// 授权名称。

	XRepoAuth string `json:"X-Repo-Auth"`
	// 组织ID。

	Namespace string `json:"namespace"`
	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。

	Project string `json:"project"`
	// 分支名称或者tag名称，如果没有提供，使用默认分支。

	Ref *string `json:"ref,omitempty"`
}

func (o ListCommitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitsRequest struct{}"
	}

	return strings.Join([]string{"ListCommitsRequest", string(data)}, " ")
}
