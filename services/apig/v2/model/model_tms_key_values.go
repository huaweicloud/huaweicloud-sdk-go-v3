package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsKeyValues struct {

	// 键。  支持可用 UTF-8 格式表示的字母(包含中文)、数字和空格，以及以下字符： _ . : = + - @； \\_sys\\_开头属于系统标签，租户不能输入
	Key *string `json:"key,omitempty"`

	// 值。  支持可用 UTF-8 格式表示的字母(包含中文)、数字和空格，以及以下字符： _ . : / = + - @
	Values *[]string `json:"values,omitempty"`
}

func (o TmsKeyValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsKeyValues struct{}"
	}

	return strings.Join([]string{"TmsKeyValues", string(data)}, " ")
}
