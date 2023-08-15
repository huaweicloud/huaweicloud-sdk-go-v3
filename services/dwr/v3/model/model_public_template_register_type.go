package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicTemplateRegisterType init_created新建，submit_approve等待审核，deprecate_approve申请禁用
type PublicTemplateRegisterType struct {
	value string
}

type PublicTemplateRegisterTypeEnum struct {
	INIT_CREATEDSUBMIT_APPROVEDEPRECATE_APPROVE PublicTemplateRegisterType
}

func GetPublicTemplateRegisterTypeEnum() PublicTemplateRegisterTypeEnum {
	return PublicTemplateRegisterTypeEnum{
		INIT_CREATEDSUBMIT_APPROVEDEPRECATE_APPROVE: PublicTemplateRegisterType{
			value: "init_created，submit_approve，deprecate_approve",
		},
	}
}

func (c PublicTemplateRegisterType) Value() string {
	return c.value
}

func (c PublicTemplateRegisterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicTemplateRegisterType) UnmarshalJSON(b []byte) error {
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
