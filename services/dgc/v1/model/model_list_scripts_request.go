package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptsRequest Request Object
type ListScriptsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 分页参数:每页限定数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数：页数
	Offset *int32 `json:"offset,omitempty"`

	// 脚本名称
	ScriptName *string `json:"scriptName,omitempty"`
}

func (o ListScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptsRequest", string(data)}, " ")
}
