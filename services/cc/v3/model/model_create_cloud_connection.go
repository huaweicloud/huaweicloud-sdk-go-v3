package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudConnection 创建云连接实例的详细信息。
type CreateCloudConnection struct {

	// 云连接实例的名字。只能由中文、英文字母、数字、下划线、中划线、点组成。
	Name string `json:"name"`

	// 云连接实例的描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 云连接实例所属的企业项目ID。企业项目账号必填；非企业项目账号不填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateCloudConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudConnection struct{}"
	}

	return strings.Join([]string{"CreateCloudConnection", string(data)}, " ")
}
