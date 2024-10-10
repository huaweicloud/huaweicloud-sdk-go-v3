package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectionReq 修改连接请求。
type UpdateConnectionReq struct {

	// 连接名称。
	Name *string `json:"name,omitempty"`

	// 数据库类型。
	DbType *string `json:"db_type,omitempty"`

	Config *ConnectionConfig `json:"config,omitempty"`

	// 连接描述。
	Description *string `json:"description,omitempty"`

	Endpoint *BaseEndpoint `json:"endpoint,omitempty"`

	Vpc *CloudVpcInfo `json:"vpc,omitempty"`

	Ssl *EndpointSslConfig `json:"ssl,omitempty"`

	Cloud *CloudBaseInfo `json:"cloud,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o UpdateConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionReq struct{}"
	}

	return strings.Join([]string{"UpdateConnectionReq", string(data)}, " ")
}
