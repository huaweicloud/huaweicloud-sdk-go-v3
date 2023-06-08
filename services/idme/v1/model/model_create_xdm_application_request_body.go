package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateXdmApplicationRequestBody struct {

	// 应用的中文名称。
	AppNameCn string `json:"app_name_cn"`

	// 应用的英文名称。
	AppNameEn string `json:"app_name_en"`

	// 应用的中文描述。
	AppDesc *string `json:"app_desc,omitempty"`

	// 应用的英文描述。
	AppDescEn *string `json:"app_desc_en,omitempty"`

	// 操作类型。
	OperateType *string `json:"operate_type,omitempty"`

	// 环境标识。 - dev：用于开发环境。 - sit：用于功能测试环境。 - uat：用于用户测试环境。 - train：用于培训环境。 - beta：用于灰度部署环境。 - production：用于生产环境。
	AppEnv string `json:"app_env"`

	// 数据库类型，支持MySQL和PostgreSQL。
	DatabaseType string `json:"database_type"`

	// 应用责任人。
	AppUserList []AppUserList `json:"app_user_list"`

	// 认证数据源中文名称。
	CertifiedDataSourceName *string `json:"certified_data_source_name,omitempty"`

	// 认证数据源编码。
	CertifiedDataSourceNumber *string `json:"certified_data_source_number,omitempty"`

	// 集成模式。 - API - SDK
	IntegrationMode string `json:"integration_mode"`

	// 元模型同步。
	MetadataSynchronization *bool `json:"metadata_synchronization,omitempty"`
}

func (o CreateXdmApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateXdmApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateXdmApplicationRequestBody", string(data)}, " ")
}
