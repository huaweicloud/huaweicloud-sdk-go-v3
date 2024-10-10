package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionResponse Response Object
type CreateConnectionResponse struct {

	// 连接ID。
	ConnectionId *string `json:"connection_id,omitempty"`

	// 连接名称。
	Name *string `json:"name,omitempty"`

	// 连接创建时间，格式为时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 连接类型。
	DbType *string `json:"db_type,omitempty"`

	Config *ConnectionConfig `json:"config,omitempty"`

	Endpoint *BaseEndpoint `json:"endpoint,omitempty"`

	Vpc *CloudVpcInfo `json:"vpc,omitempty"`

	Ssl *EndpointSslConfig `json:"ssl,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 描述。
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionResponse", string(data)}, " ")
}
