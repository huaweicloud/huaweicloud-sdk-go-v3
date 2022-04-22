package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateDependencyRequestBody struct {

	// depend_type为zip类型时必填，为文件流格式。
	DependFile *string `json:"depend_file,omitempty"`

	// depend_type为obs类型时，依赖包在obs的存储地址。
	DependLink *string `json:"depend_link,omitempty"`

	// 导入类型,目前支持obs和zip。
	DependType string `json:"depend_type"`

	// 运行时语言， Java11、Nodejs14:、Python3:在type为v2时支持。
	Runtime UpdateDependencyRequestBodyRuntime `json:"runtime"`

	// 依赖包名称。必须以大、小写字母开头，以字母或数字结尾，只能由字母、数字、下划线、点和中划线组成，长度不超过96个字符。
	Name string `json:"name"`

	// 依赖包描述，不超过512个字符。
	Description *string `json:"description,omitempty"`
}

func (o UpdateDependencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDependencyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDependencyRequestBody", string(data)}, " ")
}

type UpdateDependencyRequestBodyRuntime struct {
	value string
}

type UpdateDependencyRequestBodyRuntimeEnum struct {
	JAVA8           UpdateDependencyRequestBodyRuntime
	JAVA11          UpdateDependencyRequestBodyRuntime
	NODE_JS6_10     UpdateDependencyRequestBodyRuntime
	NODE_JS8_10     UpdateDependencyRequestBodyRuntime
	NODE_JS10_16    UpdateDependencyRequestBodyRuntime
	NODE_JS12_13    UpdateDependencyRequestBodyRuntime
	NODE_JS14_18    UpdateDependencyRequestBodyRuntime
	PYTHON2_7       UpdateDependencyRequestBodyRuntime
	PYTHON3_6       UpdateDependencyRequestBodyRuntime
	PYTHON3_9       UpdateDependencyRequestBodyRuntime
	GO1_8           UpdateDependencyRequestBodyRuntime
	GO1_X           UpdateDependencyRequestBodyRuntime
	C__NET_CORE_2_0 UpdateDependencyRequestBodyRuntime
	C__NET_CORE_2_1 UpdateDependencyRequestBodyRuntime
	C__NET_CORE_3_1 UpdateDependencyRequestBodyRuntime
	PHP7_3          UpdateDependencyRequestBodyRuntime
}

func GetUpdateDependencyRequestBodyRuntimeEnum() UpdateDependencyRequestBodyRuntimeEnum {
	return UpdateDependencyRequestBodyRuntimeEnum{
		JAVA8: UpdateDependencyRequestBodyRuntime{
			value: "Java8",
		},
		JAVA11: UpdateDependencyRequestBodyRuntime{
			value: "Java11",
		},
		NODE_JS6_10: UpdateDependencyRequestBodyRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: UpdateDependencyRequestBodyRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: UpdateDependencyRequestBodyRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: UpdateDependencyRequestBodyRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: UpdateDependencyRequestBodyRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: UpdateDependencyRequestBodyRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: UpdateDependencyRequestBodyRuntime{
			value: "Python3.6",
		},
		PYTHON3_9: UpdateDependencyRequestBodyRuntime{
			value: "Python3.9",
		},
		GO1_8: UpdateDependencyRequestBodyRuntime{
			value: "Go1.8",
		},
		GO1_X: UpdateDependencyRequestBodyRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: UpdateDependencyRequestBodyRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: UpdateDependencyRequestBodyRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: UpdateDependencyRequestBodyRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: UpdateDependencyRequestBodyRuntime{
			value: "PHP7.3",
		},
	}
}

func (c UpdateDependencyRequestBodyRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDependencyRequestBodyRuntime) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
