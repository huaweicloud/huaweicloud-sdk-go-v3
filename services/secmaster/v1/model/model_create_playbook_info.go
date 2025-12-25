package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookInfo 剧本创建参数信息
type CreatePlaybookInfo struct {

	// 剧本名称
	Name string `json:"name"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o CreatePlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookInfo struct{}"
	}

	return strings.Join([]string{"CreatePlaybookInfo", string(data)}, " ")
}
