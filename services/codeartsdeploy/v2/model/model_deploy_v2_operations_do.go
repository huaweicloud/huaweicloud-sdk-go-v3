package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeployV2OperationsDo struct {

	// 步骤id
	Id *string `json:"id,omitempty"`

	// 步骤名称
	Name *string `json:"name,omitempty"`

	// 步骤描述
	Description *string `json:"description,omitempty"`

	// 下载地址
	Code *string `json:"code,omitempty"`

	// 步骤详细定义
	Params *string `json:"params,omitempty"`

	// 入口函数
	Entrance *string `json:"entrance,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`
}

func (o DeployV2OperationsDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployV2OperationsDo struct{}"
	}

	return strings.Join([]string{"DeployV2OperationsDo", string(data)}, " ")
}
