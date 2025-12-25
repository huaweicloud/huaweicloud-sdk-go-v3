package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentConfigurationParamVo struct {

	// id
	ConfigurationId *string `json:"configuration_id,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 参数
	Param *string `json:"param,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o ComponentConfigurationParamVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentConfigurationParamVo struct{}"
	}

	return strings.Join([]string{"ComponentConfigurationParamVo", string(data)}, " ")
}
