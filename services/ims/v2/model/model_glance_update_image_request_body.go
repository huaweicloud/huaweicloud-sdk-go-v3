package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 原生更新接口请求体
type GlanceUpdateImageRequestBody struct {
	// 所需进行的更新操作的类型：替换、添加、删除。取值范围：replace、add、remove

	Op GlanceUpdateImageRequestBodyOp `json:"op"`
	// 所要操作的属性名称。 replace和remove操作取值只能是镜像当前已有的属性、add操作取值只能是镜像当前不存在的属性，需要在属性名称前加”/”

	Path string `json:"path"`
	// 所需更新/添加属性的值

	Value *string `json:"value,omitempty"`
}

func (o GlanceUpdateImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageRequestBody", string(data)}, " ")
}

type GlanceUpdateImageRequestBodyOp struct {
	value string
}

type GlanceUpdateImageRequestBodyOpEnum struct {
	REPLACE GlanceUpdateImageRequestBodyOp
	ADD     GlanceUpdateImageRequestBodyOp
	REMOVE  GlanceUpdateImageRequestBodyOp
}

func GetGlanceUpdateImageRequestBodyOpEnum() GlanceUpdateImageRequestBodyOpEnum {
	return GlanceUpdateImageRequestBodyOpEnum{
		REPLACE: GlanceUpdateImageRequestBodyOp{
			value: "replace",
		},
		ADD: GlanceUpdateImageRequestBodyOp{
			value: "add",
		},
		REMOVE: GlanceUpdateImageRequestBodyOp{
			value: "remove",
		},
	}
}

func (c GlanceUpdateImageRequestBodyOp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageRequestBodyOp) UnmarshalJSON(b []byte) error {
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
