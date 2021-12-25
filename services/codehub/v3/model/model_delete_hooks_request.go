package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteHooksRequest struct {
	// 组名

	GroupName string `json:"group_name"`
	// 通过id删除指定仓库的hook

	HookId int32 `json:"hook_id"`
	// 仓库名

	RepositoryName string `json:"repository_name"`
}

func (o DeleteHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHooksRequest struct{}"
	}

	return strings.Join([]string{"DeleteHooksRequest", string(data)}, " ")
}
