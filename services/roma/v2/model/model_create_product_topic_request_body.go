package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateProductTopicRequestBody struct {

	// 主题权限 0-发布 1-订阅
	Permission CreateProductTopicRequestBodyPermission `json:"permission"`

	// 产品级主题名称<br>Topic类格式必须以“/”进行分层，区分每个类目。其中第一个为用户自定义的版本号；第二个已经规定好，为${deviceId}通配设备ID；第三个为用户自定义的topic类名（即本字段）。Topic类组成即为：/${version}/${deviceId}/${customizePart}。简单来说，Topic类：/v1/${deviceId}/customizePart是具体Topic：/v1/deviceid1/customizePart1和/v1/deviceid2/customizePart2等的集合。
	Name string `json:"name"`

	// 版本号,输入2-50个字符。值以字母或数字开头和结尾。仅允许使用字母，数字，中划线和点号。
	Version string `json:"version"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o CreateProductTopicRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductTopicRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProductTopicRequestBody", string(data)}, " ")
}

type CreateProductTopicRequestBodyPermission struct {
	value int32
}

type CreateProductTopicRequestBodyPermissionEnum struct {
	E_0 CreateProductTopicRequestBodyPermission
	E_1 CreateProductTopicRequestBodyPermission
}

func GetCreateProductTopicRequestBodyPermissionEnum() CreateProductTopicRequestBodyPermissionEnum {
	return CreateProductTopicRequestBodyPermissionEnum{
		E_0: CreateProductTopicRequestBodyPermission{
			value: 0,
		}, E_1: CreateProductTopicRequestBodyPermission{
			value: 1,
		},
	}
}

func (c CreateProductTopicRequestBodyPermission) Value() int32 {
	return c.value
}

func (c CreateProductTopicRequestBodyPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProductTopicRequestBodyPermission) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
