package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowNamespace struct {

	// id
	Id int32 `json:"id"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Name string `json:"name"`

	// IAM用户名
	CreatorName string `json:"creator_name"`

	// 用户权限。7表示管理权限，3表示编辑权限，1表示读取权限。
	Auth ShowNamespaceAuth `json:"auth"`
}

func (o ShowNamespace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNamespace struct{}"
	}

	return strings.Join([]string{"ShowNamespace", string(data)}, " ")
}

type ShowNamespaceAuth struct {
	value int32
}

type ShowNamespaceAuthEnum struct {
	E_7 ShowNamespaceAuth
	E_3 ShowNamespaceAuth
	E_1 ShowNamespaceAuth
}

func GetShowNamespaceAuthEnum() ShowNamespaceAuthEnum {
	return ShowNamespaceAuthEnum{
		E_7: ShowNamespaceAuth{
			value: 7,
		}, E_3: ShowNamespaceAuth{
			value: 3,
		}, E_1: ShowNamespaceAuth{
			value: 1,
		},
	}
}

func (c ShowNamespaceAuth) Value() int32 {
	return c.value
}

func (c ShowNamespaceAuth) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowNamespaceAuth) UnmarshalJSON(b []byte) error {
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
