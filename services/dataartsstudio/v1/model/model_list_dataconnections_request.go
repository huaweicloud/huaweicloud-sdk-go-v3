package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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

	// 数据连接中数据库的ip，支持模糊搜索
	Ip *string `json:"ip,omitempty"`
}

func (o ListDataconnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataconnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListDataconnectionsRequest", string(data)}, " ")
}
