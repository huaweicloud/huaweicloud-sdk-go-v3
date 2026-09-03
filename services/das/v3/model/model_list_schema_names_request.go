package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemaNamesRequest Request Object
type ListSchemaNamesRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 数据库名
	DbName string `json:"db_name"`

	// 对象类型
	ObjType *string `json:"obj_type,omitempty"`

	// 是否包含所有用户
	IsWithAllUser *string `json:"is_with_all_user,omitempty"`

	// 节点类型
	NodeType *string `json:"node_type,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o ListSchemaNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemaNamesRequest struct{}"
	}

	return strings.Join([]string{"ListSchemaNamesRequest", string(data)}, " ")
}
