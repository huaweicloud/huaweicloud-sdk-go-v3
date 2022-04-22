package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项
type UpdateConfigMap struct {

	// 配置项描述,最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty"`

	// configs是一个字典，由多个键值对组成，json化后最大总长度为1048576，key和value均为字符串。键值对中key由大小写字母或中划线开头，由数字、大小写字母、点号（.）、中划线（-）、下划线（_）组成，最小长度为1，最大长度63个字符, 键值对中的value无其它限制。 注：configs字典的长度即字典转为标准的字符串后的长度，例如字典{\"a\": \"b\"}转为标准字符串后为'{\"a\": \"b\"}'，长度为10
	Configs map[string]string `json:"configs,omitempty"`
}

func (o UpdateConfigMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigMap struct{}"
	}

	return strings.Join([]string{"UpdateConfigMap", string(data)}, " ")
}
