package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TenantCmkDto struct {

	// **参数解释：** 加密主密钥的名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	CmkKeyName *string `json:"cmk_key_name,omitempty"`

	// **参数解释：** 加密主密钥的id。 **取值范围：** 字符串长度不少于1，不超过1000。
	CmkKeyId *string `json:"cmk_key_id,omitempty"`

	// **参数解释：** 加密主密钥key的状态。 **取值范围：** 1表示待激活状态,2 表示启用状态,3 表示禁用状态,4 表示计划删除状态,5 表示等待导入状态。
	KeyState *TenantCmkDtoKeyState `json:"key_state,omitempty"`
}

func (o TenantCmkDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantCmkDto struct{}"
	}

	return strings.Join([]string{"TenantCmkDto", string(data)}, " ")
}

type TenantCmkDtoKeyState struct {
	value string
}

type TenantCmkDtoKeyStateEnum struct {
	E_1 TenantCmkDtoKeyState
	E_2 TenantCmkDtoKeyState
	E_3 TenantCmkDtoKeyState
	E_4 TenantCmkDtoKeyState
	E_5 TenantCmkDtoKeyState
}

func GetTenantCmkDtoKeyStateEnum() TenantCmkDtoKeyStateEnum {
	return TenantCmkDtoKeyStateEnum{
		E_1: TenantCmkDtoKeyState{
			value: "1",
		},
		E_2: TenantCmkDtoKeyState{
			value: "2",
		},
		E_3: TenantCmkDtoKeyState{
			value: "3",
		},
		E_4: TenantCmkDtoKeyState{
			value: "4",
		},
		E_5: TenantCmkDtoKeyState{
			value: "5",
		},
	}
}

func (c TenantCmkDtoKeyState) Value() string {
	return c.value
}

func (c TenantCmkDtoKeyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TenantCmkDtoKeyState) UnmarshalJSON(b []byte) error {
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
