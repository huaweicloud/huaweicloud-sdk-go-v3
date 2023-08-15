package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateCredentialOption
type UpdateCredentialOption struct {

	// 访问密钥状态。取值为：“active”（启用）或者 “inactive”（停用）。status与description至少填写一个。
	Status *UpdateCredentialOptionStatus `json:"status,omitempty"`

	// 访问密钥描述信息。status与description至少填写一个。
	Description *string `json:"description,omitempty"`
}

func (o UpdateCredentialOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCredentialOption struct{}"
	}

	return strings.Join([]string{"UpdateCredentialOption", string(data)}, " ")
}

type UpdateCredentialOptionStatus struct {
	value string
}

type UpdateCredentialOptionStatusEnum struct {
	ACTIVE   UpdateCredentialOptionStatus
	INACTIVE UpdateCredentialOptionStatus
}

func GetUpdateCredentialOptionStatusEnum() UpdateCredentialOptionStatusEnum {
	return UpdateCredentialOptionStatusEnum{
		ACTIVE: UpdateCredentialOptionStatus{
			value: "active",
		},
		INACTIVE: UpdateCredentialOptionStatus{
			value: "inactive",
		},
	}
}

func (c UpdateCredentialOptionStatus) Value() string {
	return c.value
}

func (c UpdateCredentialOptionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCredentialOptionStatus) UnmarshalJSON(b []byte) error {
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
