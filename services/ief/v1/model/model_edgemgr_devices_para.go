package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgemgrDevicesPara 终端设备属性
type EdgemgrDevicesPara struct {

	// 终端设备描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty"`

	// 终端设备静态属性
	Attributes map[string]ValueInAttributes `json:"attributes,omitempty"`
}

func (o EdgemgrDevicesPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgemgrDevicesPara struct{}"
	}

	return strings.Join([]string{"EdgemgrDevicesPara", string(data)}, " ")
}
