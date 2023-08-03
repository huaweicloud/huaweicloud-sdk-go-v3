package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataConnectorRequest Request Object
type ListDataConnectorRequest struct {

	// 连接id
	ConnectorId *string `json:"connector_id,omitempty"`

	// 数据源类别
	SourceType *string `json:"source_type,omitempty"`

	// 数据连接名称
	ConnectorName *string `json:"connector_name,omitempty"`

	// 每页返回的资源个数
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 数据连接是否有效
	Available *bool `json:"available,omitempty"`
}

func (o ListDataConnectorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataConnectorRequest struct{}"
	}

	return strings.Join([]string{"ListDataConnectorRequest", string(data)}, " ")
}
