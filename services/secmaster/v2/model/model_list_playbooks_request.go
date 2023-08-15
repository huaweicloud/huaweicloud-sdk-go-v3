package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybooksRequest Request Object
type ListPlaybooksRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// 搜索关键字
	SearchTxt *string `json:"search_txt,omitempty"`

	// component id.
	ComponentId *string `json:"component_id,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// request offset, from 0
	Offset *int32 `json:"offset,omitempty"`

	// request limit size
	Limit *int32 `json:"limit,omitempty"`

	// 剧本描述
	Description *string `json:"description,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 剧本名称
	Name *string `json:"name,omitempty"`
}

func (o ListPlaybooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybooksRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybooksRequest", string(data)}, " ")
}
