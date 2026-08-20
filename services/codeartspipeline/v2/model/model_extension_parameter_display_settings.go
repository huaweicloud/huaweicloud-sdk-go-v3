package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionParameterDisplaySettings UI控件配置
type ExtensionParameterDisplaySettings struct {

	// 控件类型，如 Select/CodeText/Radio/SingleLineText/Hidden
	DevCloudControlType *string `json:"DevCloud.ControlType,omitempty"`

	// 默认选中值。可能为字符串，也可能为对象(如 {displayName, value})。
	DevCloudControlTypeDefault *string `json:"DevCloud.ControlType.Default,omitempty"`

	// 下拉选项(Select类型)。
	DevCloudControlTypeSelect *[]string `json:"DevCloud.ControlType.Select,omitempty"`

	// 单选选项(Radio类型)。
	DevCloudControlTypeRadio *[]ExtensionRadioOption `json:"DevCloud.ControlType.Radio,omitempty"`
}

func (o ExtensionParameterDisplaySettings) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionParameterDisplaySettings struct{}"
	}

	return strings.Join([]string{"ExtensionParameterDisplaySettings", string(data)}, " ")
}
