package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigsMap struct {

	// 配置项的名称，以英文字母或中划线开头，由英文字母、数字、点号、中划线和下划线组成，长度1到63个字符。
	Key string `json:"key"`

	// 配置项的属性名，以英文小写字母开头，由中文字符，英文字母，数字，下划线和中划线组成，不能以中划线结尾，长度4-64位。
	Name string `json:"name"`
}

func (o ConfigsMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigsMap struct{}"
	}

	return strings.Join([]string{"ConfigsMap", string(data)}, " ")
}
