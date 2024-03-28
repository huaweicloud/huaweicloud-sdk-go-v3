package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseRequestBody 创建数据库的请求body体。
type CreateDatabaseRequestBody struct {

	// 新增数据库名称。 说明： “default”为内置数据库，不能创建名为“default”的数据库。
	DatabaseName string `json:"database_name"`

	// 新增数据库的描述信息。
	Description *string `json:"description,omitempty"`

	// 企业项目ID，“0”表示default，即默认的企业项目。关于如何设置企业项目请参考《企业管理用户指南》。 说明： 开通了企业管理服务的用户可设置该参数绑定指定的项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateDatabaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequestBody", string(data)}, " ")
}
