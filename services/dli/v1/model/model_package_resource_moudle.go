package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackageResourceMoudle 系统内置资源模块列表。
type PackageResourceMoudle struct {

	// 模块名。
	ModuleName *string `json:"module_name,omitempty"`

	// 模块类型。
	ModuleType *string `json:"module_type,omitempty"`

	// \"UPLOADING\"表示正在上传 \"READY\"表示模块包已上传 \"FAILED\"表示模块包上传失败
	Status *string `json:"status,omitempty"`

	// 模块描述。
	Description *string `json:"description,omitempty"`

	// 该模块包含的资源包名列表。
	Resources *[]string `json:"resources,omitempty"`

	// 模块上传的unix时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 模块更新的unix时间。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o PackageResourceMoudle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageResourceMoudle struct{}"
	}

	return strings.Join([]string{"PackageResourceMoudle", string(data)}, " ")
}
