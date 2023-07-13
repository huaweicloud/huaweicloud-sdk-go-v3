package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchDwByTypeRequest Request Object
type SearchDwByTypeRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 是否查询最新的
	ForceRefresh *bool `json:"force_refresh,omitempty"`

	// 数据连接类型
	DwType string `json:"dw_type"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// limit
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchDwByTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDwByTypeRequest struct{}"
	}

	return strings.Join([]string{"SearchDwByTypeRequest", string(data)}, " ")
}
