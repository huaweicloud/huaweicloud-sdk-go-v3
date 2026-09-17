package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionRequest Request Object
type ListConnectionRequest struct {

	// 数据库实例地址/实例名称/备注等关键字
	Condition *string `json:"condition,omitempty"`

	// 每页记录数
	Perpage *string `json:"perpage,omitempty"`

	// 页码
	Curpage *string `json:"curpage,omitempty"`

	// 数据库来源类型
	NetworkType *string `json:"network_type,omitempty"`

	// 数据库引擎类型
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 连接类型
	ConnectionType *string `json:"connection_type,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionRequest", string(data)}, " ")
}
