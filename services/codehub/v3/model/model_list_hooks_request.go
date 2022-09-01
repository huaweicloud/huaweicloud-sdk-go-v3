package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHooksRequest struct {

	// 组名
	GroupName string `json:"group_name" xml:"group_name"`

	// hook id
	HookId *string `json:"hook_id,omitempty" xml:"hook_id"`

	// 仓库名
	RepositoryName string `json:"repository_name" xml:"repository_name"`
}

func (o ListHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHooksRequest struct{}"
	}

	return strings.Join([]string{"ListHooksRequest", string(data)}, " ")
}
