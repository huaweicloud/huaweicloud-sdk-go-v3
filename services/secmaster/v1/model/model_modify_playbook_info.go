package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPlaybookInfo 修改剧本详情
type ModifyPlaybookInfo struct {

	// 剧本名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 启用的剧本版本ID
	ActiveVersionId *string `json:"active_version_id,omitempty"`
}

func (o ModifyPlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPlaybookInfo struct{}"
	}

	return strings.Join([]string{"ModifyPlaybookInfo", string(data)}, " ")
}
