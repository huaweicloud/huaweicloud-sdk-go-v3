package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEnvironmentVariableV2Response struct {

	// 变量值支持英文字母、数字、英文格式的下划线、中划线，斜线（/）、点、冒号，1 ~ 255个字符。
	VariableValue string `json:"variable_value" xml:"variable_value"`

	// 环境编号
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 变量名，支持英文字母、数字、英文格式的下划线、中划线，必须以英文字母开头，3~32个字符。在API定义中等于#Name的值#部分（区分大小写），发布到环境里的API被变量值换。 > 中文字符必须为UTF-8或者unicode编码。
	VariableName *string `json:"variable_name,omitempty" xml:"variable_name"`

	// 环境变量编号
	Id             *string `json:"id,omitempty" xml:"id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEnvironmentVariableV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentVariableV2Response struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentVariableV2Response", string(data)}, " ")
}
