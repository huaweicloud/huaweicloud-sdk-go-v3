package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagRequest Request Object
type CreateTagRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth"`

	// 命名空间ID或者URL编码名称。
	Namespace string `json:"namespace"`

	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。
	Project string `json:"project"`

	// 分支名称或者tag标签名称或者commit sha。
	Ref string `json:"ref"`

	Body *TagCreate `json:"body,omitempty"`
}

func (o CreateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTagRequest", string(data)}, " ")
}
