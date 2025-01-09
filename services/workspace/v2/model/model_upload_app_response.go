package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAppResponse Response Object
type UploadAppResponse struct {

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 租户id。
	TenantId *string `json:"tenant_id,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	AuthorizationType *AssignType `json:"authorization_type,omitempty"`

	AppFileStore *FileStoreLink `json:"app_file_store,omitempty"`

	// 应用图标路径。
	AppIconUrl *string `json:"app_icon_url,omitempty"`

	InstallType *InstallType `json:"install_type,omitempty"`

	// 安装命令(静默安装命令)。 例: ${FILE_PATH} /S。 预定义变量将采用以下值: ${FILE_PATH}: 应用安装包在桌面本地的存储路径。
	InstallCommand *string `json:"install_command,omitempty"`

	// 卸载命令(静默卸载命令)。 例: msiexec /uninstall ${FILE_PATH} /quiet。 预定义变量将采用以下值: ${FILE_PATH}: 应用安装包在桌面本地的存储路径。
	UninstallCommand *string `json:"uninstall_command,omitempty"`

	SupportOs *OsTypeEnum `json:"support_os,omitempty"`

	Status *AppStatusEnum `json:"status,omitempty"`

	ApplicationSource *AppSourceType `json:"application_source,omitempty"`

	// 应用创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 分类ID。
	CatalogId *string `json:"catalog_id,omitempty"`

	// 分类名称。
	Catalog *string `json:"catalog,omitempty"`

	// 安装信息。
	InstallInfo    *string `json:"install_info,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAppResponse struct{}"
	}

	return strings.Join([]string{"UploadAppResponse", string(data)}, " ")
}
