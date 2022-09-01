package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddHooksRequest struct {

	// 组名
	GroupName string `json:"group_name" xml:"group_name"`

	// 仓库名
	RepositoryName string `json:"repository_name" xml:"repository_name"`

	Body *RepositoryHookRequest `json:"body,omitempty" xml:"body"`
}

func (o AddHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHooksRequest struct{}"
	}

	return strings.Join([]string{"AddHooksRequest", string(data)}, " ")
}
