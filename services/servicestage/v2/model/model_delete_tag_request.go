package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteTagRequest struct {
	// 授权名称。

	XRepoAuth string `json:"X-Repo-Auth"`
	// 组织ID。

	Namespace string `json:"namespace"`
	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。

	Project string `json:"project"`
	// tag标签名称。

	TagName string `json:"tag_name"`
}

func (o DeleteTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagRequest", string(data)}, " ")
}
