package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantVersionConfigResponse Response Object
type ShowTenantVersionConfigResponse struct {

	// 租户版本配置ID
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 关联版本配置表ID
	VersionConfigId *string `json:"version_config_id,omitempty"`

	// 自定义版本说明
	CustomReleaseNote *string `json:"custom_release_note,omitempty"`

	// 更新说明
	ReleaseNote *string `json:"release_note,omitempty"`

	// 版本下载地址
	VersionDownloadUrl *string `json:"version_download_url,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTenantVersionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantVersionConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantVersionConfigResponse", string(data)}, " ")
}
