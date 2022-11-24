package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateArgs struct {

	// 模板参数名。 目前仅支持sub_customer_name：表明企业主创建企业子的名字
	Key string `json:"key"`

	// 模板参数值。 key对应的取值。
	Value string `json:"value"`
}

func (o TemplateArgs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateArgs struct{}"
	}

	return strings.Join([]string{"TemplateArgs", string(data)}, " ")
}
