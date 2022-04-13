package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddHooksRequest struct {
	// 组名

	GroupName string `json:"group_name"`
	// 仓库名

	RepositoryName string `json:"repository_name"`

	Body *RepositoryHookRequest `json:"body,omitempty"`
}

func (o AddHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHooksRequest struct{}"
	}

	return strings.Join([]string{"AddHooksRequest", string(data)}, " ")
}
