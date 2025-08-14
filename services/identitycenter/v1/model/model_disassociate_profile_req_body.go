package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisassociateProfileReqBody 解除与用户或组绑定的所有账号授权关联的请求体
type DisassociateProfileReqBody struct {

	// 用户或用户组的唯一标识ID
	Id string `json:"id"`

	// 解除绑定的实体类型，可为用户或者用户组
	AccessorType DisassociateProfileReqBodyAccessorType `json:"accessor_type"`

	// 关联到IAM身份中心实例的身份源的全局唯一标识符（ID）。
	IdentityStoreId string `json:"identity_store_id"`
}

func (o DisassociateProfileReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateProfileReqBody struct{}"
	}

	return strings.Join([]string{"DisassociateProfileReqBody", string(data)}, " ")
}

type DisassociateProfileReqBodyAccessorType struct {
	value string
}

type DisassociateProfileReqBodyAccessorTypeEnum struct {
	GROUP DisassociateProfileReqBodyAccessorType
	USER  DisassociateProfileReqBodyAccessorType
}

func GetDisassociateProfileReqBodyAccessorTypeEnum() DisassociateProfileReqBodyAccessorTypeEnum {
	return DisassociateProfileReqBodyAccessorTypeEnum{
		GROUP: DisassociateProfileReqBodyAccessorType{
			value: "GROUP",
		},
		USER: DisassociateProfileReqBodyAccessorType{
			value: "USER",
		},
	}
}

func (c DisassociateProfileReqBodyAccessorType) Value() string {
	return c.value
}

func (c DisassociateProfileReqBodyAccessorType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisassociateProfileReqBodyAccessorType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
