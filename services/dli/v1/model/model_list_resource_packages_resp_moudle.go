package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcePackagesRespMoudle 系统内置资源模块列表。
type ListResourcePackagesRespMoudle struct {

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

func (o ListResourcePackagesRespMoudle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcePackagesRespMoudle struct{}"
	}

	return strings.Join([]string{"ListResourcePackagesRespMoudle", string(data)}, " ")
}
