package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudConnection 创建云连接实例的详细信息。
type CreateCloudConnection struct {

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateCloudConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudConnection struct{}"
	}

	return strings.Join([]string{"CreateCloudConnection", string(data)}, " ")
}
