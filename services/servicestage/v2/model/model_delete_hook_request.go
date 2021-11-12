package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteHookRequest struct {
	// 授权名称。

	XRepoAuth string `json:"X-Repo-Auth"`
	// 组织ID。

	Namespace string `json:"namespace"`
	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。

	Project string `json:"project"`
	// hook ID。

	HookId string `json:"hook_id"`
}

func (o DeleteHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHookRequest struct{}"
	}

	return strings.Join([]string{"DeleteHookRequest", string(data)}, " ")
}
