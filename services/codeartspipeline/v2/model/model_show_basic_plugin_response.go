package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBasicPluginResponse Response Object
type ShowBasicPluginResponse struct {

	// 类型
	Type *string `json:"type,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 展示名
	FriendlyName *string `json:"friendly_name,omitempty"`

	// 业务类型
	Category *string `json:"category,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 版本说明
	VersionDescription *string `json:"version_description,omitempty"`

	// 输入信息
	Inputs *[]NewExtensionInputs `json:"inputs,omitempty"`

	// 输出信息
	Outputs *[]NewExtensionOutputs `json:"outputs,omitempty"`

	Execution      *NewExtensionExecution `json:"execution,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowBasicPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBasicPluginResponse struct{}"
	}

	return strings.Join([]string{"ShowBasicPluginResponse", string(data)}, " ")
}
