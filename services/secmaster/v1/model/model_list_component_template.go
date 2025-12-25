package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListComponentTemplate struct {

	// 组件id.
	ComponentId *string `json:"component_id,omitempty"`

	// 组件名称
	ComponentName *string `json:"component_name,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	FileType *FileType `json:"file_type,omitempty"`

	// 参数
	Param *string `json:"param,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o ListComponentTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentTemplate struct{}"
	}

	return strings.Join([]string{"ListComponentTemplate", string(data)}, " ")
}
