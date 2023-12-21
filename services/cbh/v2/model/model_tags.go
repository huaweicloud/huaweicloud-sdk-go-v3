package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tags 标签信息。  > 说明： - 不允许重复key - 每实例最多允许20个标签（系统可配置），不包含_sys开头系统标签  - 优先按照创建时间的倒序排序。如果没有创建时间则按照资源名称ASCII码进行排序
type Tags struct {

	// 键。  > 说明： - key不能为空，长度1~128个字符（中文也可以输入128个字符） - 可用 UTF-8 格式表示的字母(包含中文)、数字和空格，以及以下字符： _ . : = + - @ - _sys_开头属于系统标签，租户不能输入 - 建议正则：^((?!_sys_)[\\\\p{L}\\\\p{Z}\\\\p{N}_.:=+\\\\-@]*)$
	Key string `json:"key"`

	// 值列表。  > 说明： - 长度0~255个字符（中文也可以输入255个字符） - 可用 UTF-8 格式表示的字母(包含中文)、数字和空格，以及以下字符： _ . : / = + - @ 建议正则：^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$ - 资源标签值可以为空（empty or null） - 预定义标签值不可以为空
	Values []string `json:"values"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
