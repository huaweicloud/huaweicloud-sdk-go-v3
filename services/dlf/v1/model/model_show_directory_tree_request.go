package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDirectoryTreeRequest Request Object
type ShowDirectoryTreeRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 查询的目录节点类型.
	Category *string `json:"category,omitempty"`

	// 目录编号.
	DirectoryId *string `json:"directory_id,omitempty"`

	// 目标元素名称.
	Name *string `json:"name,omitempty"`

	// 查询偏移量.
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量.
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDirectoryTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirectoryTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowDirectoryTreeRequest", string(data)}, " ")
}
