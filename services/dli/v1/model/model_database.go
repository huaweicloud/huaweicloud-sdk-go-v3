package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Database 查询所有数据库的响应参数。
type Database struct {

	// 数据库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据库的创建者。
	Owner *string `json:"owner,omitempty"`

	// 数据库中表的个数。
	TableNumber *int32 `json:"table_number,omitempty"`

	// 数据库相关的描述信息。
	Description *string `json:"description,omitempty"`

	// 企业项目ID，“0”表示default，即默认的企业项目。关于如何设置企业项目请参考《企业管理用户指南》。 说明： 开通了企业管理服务的用户可设置该参数绑定指定的项目。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 资源ID。
	ResourceId string `json:"resource_id"`
}

func (o Database) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Database struct{}"
	}

	return strings.Join([]string{"Database", string(data)}, " ")
}
