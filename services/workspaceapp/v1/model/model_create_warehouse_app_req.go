package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWarehouseAppReq 先上传应用文件，再提交向应用仓库中添加应用的请求。
type CreateWarehouseAppReq struct {

	// 应用名称,名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成。 2. 长度范围1~64个字符。
	AppName string `json:"app_name"`

	AppCategory *AppCategoryEnum `json:"app_category"`

	OsType *OsTypeEnum `json:"os_type"`

	// 版本号,名称需满足如下规则: 1. 由可见字符组成。 2. 长度范围1~64个字符。
	VersionId string `json:"version_id"`

	// 应用仓库中的应用描述。
	AppDescription *string `json:"app_description,omitempty"`

	// 版本描述,名称需满足如下规则: 1. 由可见字符组成。 2. 长度范围1~255个字符。
	VersionName string `json:"version_name"`

	// 应用在obs桶的存储路径。
	AppfileStorePath string `json:"appfile_store_path"`

	// > - 图片的默认大小当前限制为8KB，即1024 * 8字节。 > - 如果数据格式为data;image/png;base64,iVBORw0KGgoAAAANS时实际大小约为字段约为：size * 4/3 + 4bytes。
	AppIcon *string `json:"app_icon,omitempty"`

	// 应用文件大小，单位为KB。
	AppFileSize *int32 `json:"app_file_size,omitempty"`

	AppExtendedInfo *AppExtendedInfo `json:"app_extended_info,omitempty"`
}

func (o CreateWarehouseAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWarehouseAppReq struct{}"
	}

	return strings.Join([]string{"CreateWarehouseAppReq", string(data)}, " ")
}
