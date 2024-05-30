package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesResultDataValue value，统一的返回结果的外层数据结构。
type ListWorkspacesResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// WorkspaceVO信息。
	Records *[]WorkspaceVo `json:"records,omitempty"`
}

func (o ListWorkspacesResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesResultDataValue struct{}"
	}

	return strings.Join([]string{"ListWorkspacesResultDataValue", string(data)}, " ")
}
