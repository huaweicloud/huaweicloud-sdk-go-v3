package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListGaussMySqlInstancesRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。 “*”为系统保留字符，如果id是以 “*”起始，表示按照 “*”后面的值模糊匹配，否则，按照id精确匹配查询。不能只传入 “*”。

	Id *string `json:"id,omitempty"`
	// 实例名称。  “*”为系统保留字符，如果name是以 “*”起始，表示按照 “*”后面的值模糊匹配，否则，按照name精确匹配查询。不能只传入 “*”。

	Name *string `json:"name,omitempty"`
	// 按照实例类型查询。目前仅支持Cluster。

	Type *string `json:"type,omitempty"`
	// 数据库类型，现在只支持gaussdb-mysql。

	DatastoreType *string `json:"datastore_type,omitempty"`
	// 虚拟私有云ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 子网的网络ID信息。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGaussMySqlInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlInstancesRequest", string(data)}, " ")
}
