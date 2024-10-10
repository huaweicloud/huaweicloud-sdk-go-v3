package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionResp 连接响应体。
type ConnectionResp struct {

	// 连接ID。
	ConnectionId string `json:"connection_id"`

	// 连接名称。
	Name string `json:"name"`

	// 连接创建时间，格式为时间戳。
	CreateTime int64 `json:"create_time"`

	// 连接类型。
	DbType string `json:"db_type"`

	Config *ConnectionConfig `json:"config,omitempty"`

	Endpoint *BaseEndpoint `json:"endpoint"`

	Vpc *CloudVpcInfo `json:"vpc,omitempty"`

	Ssl *EndpointSslConfig `json:"ssl,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o ConnectionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionResp struct{}"
	}

	return strings.Join([]string{"ConnectionResp", string(data)}, " ")
}
