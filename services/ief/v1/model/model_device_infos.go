package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 关联设备信息
type DeviceInfos struct {
	// 设备和节点关系的名称，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64

	Relation string `json:"relation"`
	// 设备和节点关系的描述，最大长度64，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\

	Comment *string `json:"comment,omitempty"`
	// 设备ID列表

	DeviceIds []string `json:"device_ids"`
}

func (o DeviceInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceInfos struct{}"
	}

	return strings.Join([]string{"DeviceInfos", string(data)}, " ")
}
