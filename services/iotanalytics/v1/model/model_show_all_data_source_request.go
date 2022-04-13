package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAllDataSourceRequest struct {
	// filter

	Name *string `json:"name,omitempty"`
	// 数据源类型, 包括：OBS、DIS、IOTDA、SMN、FUNCTION_GRAPH、MODEL_ARTS、DCS、KAFKA、API

	Type *string `json:"type,omitempty"`
	// 每次查询返回元素的上限

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，表示从此偏移量开始查询，offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowAllDataSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllDataSourceRequest struct{}"
	}

	return strings.Join([]string{"ShowAllDataSourceRequest", string(data)}, " ")
}
