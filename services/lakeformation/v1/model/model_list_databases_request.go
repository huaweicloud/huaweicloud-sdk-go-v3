package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabasesRequest Request Object
type ListDatabasesRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如:2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// catalog名称。只能包含字母、数字和下划线,且长度为1~256个字符。
	CatalogName string `json:"catalog_name"`

	// 数据库名称通配符。只能包含中文、字母、数字和_|*.-特殊字符,且长度为1~128个字符。
	DatabaseNamePattern *string `json:"database_name_pattern,omitempty"`

	// 查询返回条数。默认值为1000。最小值为0,最大值为1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID。最小长度为0,最大长度为256。
	Marker *string `json:"marker,omitempty"`

	// 是否倒序查询。
	ReversePage *bool `json:"reverse_page,omitempty"`

	// 用户端数据库id,创建时指定,不可修改。
	ExternalDatabaseId *string `json:"external_database_id,omitempty"`

	// 是否查询被删除元数据。
	Deleted *bool `json:"deleted,omitempty"`
}

func (o ListDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabasesRequest", string(data)}, " ")
}
