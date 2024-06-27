package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateChDatabaseConfigRequestBody 修改数据同步
type UpdateChDatabaseConfigRequestBody struct {

	// 源实例ID，严格匹配UUID规则。
	SourceInstanceId string `json:"source_instance_id"`

	// 源实例节点ID。
	SourceNodeId *string `json:"source_node_id,omitempty"`

	// 数据库名。
	Database string `json:"database"`

	// 配置值。仅支持修改同步范围，不支持修改白名单或黑名单类型。  例如：创建了白名单数据同步，调用本接口修改时， 支持 \"value\"：\"{\\\"white_list\\\":\\\"test1,test2,test3\\\"}\" 不支持 \"value\"：\"{\\\"black_list\\\":\\\"test1,test2,test3\\\"}\"
	Value string `json:"value"`
}

func (o UpdateChDatabaseConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChDatabaseConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateChDatabaseConfigRequestBody", string(data)}, " ")
}
