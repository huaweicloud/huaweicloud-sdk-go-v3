package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRetrieveScriptsRequest Request Object
type ListRetrieveScriptsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId *string `json:"table_id,omitempty"`

	// 脚本名称
	ScriptName *string `json:"script_name,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释：** 查询数据限制 **取值范围：** 0-1000 **默认取值：** 不涉及
	Limit *int64 `json:"limit,omitempty"`

	// 按照属性排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序，支持 `ASC` 或 `DESC`。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListRetrieveScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetrieveScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListRetrieveScriptsRequest", string(data)}, " ")
}
