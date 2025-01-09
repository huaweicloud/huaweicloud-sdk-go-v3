package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAppReq 添加应用。
type UploadAppReq struct {

	// 应用名称,名称需满足如下规则: 1. 不可为全空格。 2. 不允许包含如下字符:^;|~`{}[]<>。 3. 长度1~128个字符。
	Name string `json:"name"`

	// 版本号。
	Version string `json:"version"`

	// 描述。
	Description string `json:"description"`

	AuthorizationType *AssignType `json:"authorization_type"`

	AppFileStore *FileStoreLink `json:"app_file_store"`

	// 图片的路径,支持使用可访问的URL地址或DataURIscheme。 * `可访问的URL` - https://xxx.x.xx.x/xxx/xx.jpg。 * `DataURIscheme` -  data;image/png;base64,iVBORw0KGgoAAAANS; 注意使用dataURLStream时，字符串最大长度为87500，即最多使用约64KB大小的图片。
	AppIconUrl *string `json:"app_icon_url,omitempty"`

	InstallType *InstallType `json:"install_type"`

	// 安装命令(静默安装命令)。 例: ${FILE_PATH} /S。 预定义变量将采用以下值: ${FILE_PATH}: 应用安装包在桌面本地的存储路径。
	InstallCommand *string `json:"install_command,omitempty"`

	// 卸载命令(静默卸载命令)。 例: msiexec /uninstall ${FILE_PATH} /quiet。 预定义变量将采用以下值: ${FILE_PATH}: 应用安装包在桌面本地的存储路径。
	UninstallCommand *string `json:"uninstall_command,omitempty"`

	SupportOs *OsTypeEnum `json:"support_os"`

	// 分类ID。
	CatalogId string `json:"catalog_id"`

	// 安装信息。
	InstallInfo *string `json:"install_info,omitempty"`
}

func (o UploadAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAppReq struct{}"
	}

	return strings.Join([]string{"UploadAppReq", string(data)}, " ")
}
