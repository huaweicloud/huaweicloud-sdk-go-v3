package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VarsStructure HCL支持参数，即，同一个模板可以给予不同的参数而达到不同的效果。  var_structure可以允许客户提交最简单的字符串类型的参数  资源编排服务支持vars_structure，vars_body和vars_uri，如果以上三种方式中声名了同一个变量，将报错400  vars_structure中的值只支持简单的字符串类型，如果需要使用其他类型，需要用户自己在HCL引用时转换， 或者用户可以使用vars_uri、vars_body，vars_uri和vars_body中支持HCL支持的各种类型以及复杂结构  如果vars_structure过大，可以使用vars_uri  注意：vars中不应该传递任何敏感信息，资源编排服务会直接明文使用、log、展示、存储对应的vars
type VarsStructure struct {

	// 参数的名字
	VarKey string `json:"var_key"`

	// 参数的值。  注意，参数需要以字符串形式存在，如果是数字，也需要以字符串形式存在，如'10'。  如果需要支持不同类型，或者复杂结构，请使用vars_uri或vars_body
	VarValue string `json:"var_value"`

	Encryption *EncryptionStructure `json:"encryption,omitempty"`
}

func (o VarsStructure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VarsStructure struct{}"
	}

	return strings.Join([]string{"VarsStructure", string(data)}, " ")
}
