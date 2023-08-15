package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataconnectionsRequest Request Object
type ListDataconnectionsRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// 数据连接名称
	Name *string `json:"name,omitempty"`

	// 数据连接类型,有HIVE,MYSQL,ORALCLE,DWS,HBASE等。
	Type *string `json:"type,omitempty"`

	// 数据条数限制
	Limit *string `json:"limit,omitempty"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`
}

func (o ListDataconnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataconnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListDataconnectionsRequest", string(data)}, " ")
}
