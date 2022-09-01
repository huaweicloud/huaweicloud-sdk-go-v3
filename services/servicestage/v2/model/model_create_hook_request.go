package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateHookRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth" xml:"X-Repo-Auth"`

	// 组织ID。
	Namespace string `json:"namespace" xml:"namespace"`

	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。
	Project string `json:"project" xml:"project"`

	Body *HookCreate `json:"body,omitempty" xml:"body"`
}

func (o CreateHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHookRequest struct{}"
	}

	return strings.Join([]string{"CreateHookRequest", string(data)}, " ")
}
