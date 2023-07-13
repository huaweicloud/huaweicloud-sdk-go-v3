package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListColumnsRequest Request Object
type ListColumnsRequest struct {

	// 数据所在空间的id值
	Workspace string `json:"workspace"`

	// 数据连接id
	ConnectionId string `json:"connection_id"`

	// 表id
	TableId string `json:"table_id"`

	// 数据条数限制
	Limit *string `json:"limit,omitempty"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`
}

func (o ListColumnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListColumnsRequest struct{}"
	}

	return strings.Join([]string{"ListColumnsRequest", string(data)}, " ")
}
