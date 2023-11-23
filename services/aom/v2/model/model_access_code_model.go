package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessCodeModel struct {

	// access_code
	AccessCode *string `json:"access_code,omitempty"`

	// access_code_id
	AccessCodeId *string `json:"access_code_id,omitempty"`

	// 创建时间
	CreateAt *int64 `json:"create_at,omitempty"`

	// 状态 enable:开启 unenable:关闭
	Status *AccessCodeModelStatus `json:"status,omitempty"`
}

func (o AccessCodeModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessCodeModel struct{}"
	}

	return strings.Join([]string{"AccessCodeModel", string(data)}, " ")
}

type AccessCodeModelStatus struct {
	value string
}

type AccessCodeModelStatusEnum struct {
	ENABLE   AccessCodeModelStatus
	UNENABLE AccessCodeModelStatus
}

func GetAccessCodeModelStatusEnum() AccessCodeModelStatusEnum {
	return AccessCodeModelStatusEnum{
		ENABLE: AccessCodeModelStatus{
			value: "enable",
		},
		UNENABLE: AccessCodeModelStatus{
			value: "unenable",
		},
	}
}

func (c AccessCodeModelStatus) Value() string {
	return c.value
}

func (c AccessCodeModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessCodeModelStatus) UnmarshalJSON(b []byte) error {
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
