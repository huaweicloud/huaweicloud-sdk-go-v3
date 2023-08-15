package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NamespaceVo struct {
	CreatedDate *sdktime.SdkTime `json:"created_date,omitempty"`

	DataSyncStatistics *DataSyncStatistics `json:"data_sync_statistics,omitempty"`

	DatasourceStatistics *DatasourceStatistics `json:"datasource_statistics,omitempty"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Id *string `json:"id,omitempty"`

	IsUsed *bool `json:"is_used,omitempty"`

	MultiActiveZone *[]MultiActiveZoneVo `json:"multi_active_zone,omitempty"`

	Name *string `json:"name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	Type *int32 `json:"type,omitempty"`

	UpdatedDate *sdktime.SdkTime `json:"updated_date,omitempty"`

	UserId *string `json:"user_id,omitempty"`
}

func (o NamespaceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceVo struct{}"
	}

	return strings.Join([]string{"NamespaceVo", string(data)}, " ")
}
