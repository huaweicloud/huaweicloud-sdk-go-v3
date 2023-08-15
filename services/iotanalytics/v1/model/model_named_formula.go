package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamedFormula 带名称的公式
type NamedFormula struct {

	// 公式名称，不能和输入参数名重复，正则：\"^[A-Za-z][A-Za-z_]{0,31}$\"
	Name string `json:"name"`

	// 公式，最多1024个字符
	Formula string `json:"formula"`
}

func (o NamedFormula) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamedFormula struct{}"
	}

	return strings.Join([]string{"NamedFormula", string(data)}, " ")
}
