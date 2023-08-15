package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppOutputParameterDto 应用输出参数
type AppOutputParameterDto struct {

	// 参数名称，单个应用内唯一。取值范围：长度为[1,32]，以小写字母开头，允许出现中划线(-)、小写字母和数字，且必须以小写字母或数字结尾。
	Name string `json:"name"`

	// 参数描述。取值范围：[0-255]
	Description *string `json:"description,omitempty"`

	// 参数是否必须
	Required bool `json:"required"`

	// 参数类型。取值：[STRING，FILE，DIRECTORY，ENUM]
	Type string `json:"type"`

	// 提示用户参数填写的格式，取值范围：[0-64]。对于STRING类型，匹配字符串内容，比如后缀约束.fastq；对于ENUM类型，可以提示一定要在param_enum列表范围内取值；对于FILE类型，约束文件后缀类型；对于DIRECTORY类型，提示目录下需要包含哪些文件；
	Pattern *string `json:"pattern,omitempty"`

	// 参数取值 如填写，只支持填一项，根据参数类型进行不同的校验
	Values *[]string `json:"values,omitempty"`

	// 枚举参数的取值列表，列表最大长度20，每一项字符串最长128。参数类型为ENUM时需要填此字段
	Enum *[]string `json:"enum,omitempty"`
}

func (o AppOutputParameterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppOutputParameterDto struct{}"
	}

	return strings.Join([]string{"AppOutputParameterDto", string(data)}, " ")
}
