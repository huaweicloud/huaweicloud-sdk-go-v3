package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHookRequest Request Object
type CreateHookRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth"`

	// 命名空间ID或者URL编码名称。
	Namespace string `json:"namespace"`

	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。
	Project string `json:"project"`

	Body *HookCreate `json:"body,omitempty"`
}

func (o CreateHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHookRequest struct{}"
	}

	return strings.Join([]string{"CreateHookRequest", string(data)}, " ")
}
