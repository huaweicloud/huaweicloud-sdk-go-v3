package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateTenantRepoEncryptionSettingResponse Response Object
type UpdateTenantRepoEncryptionSettingResponse struct {

	// **参数解释：** 租户id。 **取值范围：** 字符串长度不少于1，不超过1000。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释：** 加密类型。 **取值范围：** KMS表示开启KMS加密，normal或者null表示未开启KMS加密。
	EncryptionType *string `json:"encryption_type,omitempty"`

	// **参数解释：** 是否开启租户下默认加密设置。
	DefaultEncryptionEnabled *bool `json:"default_encryption_enabled,omitempty"`

	// **参数解释：** 加密主密钥的名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	CmkKeyName *string `json:"cmk_key_name,omitempty"`

	// **参数解释：** 加密主密钥的id。 **取值范围：** 字符串长度不少于1，不超过1000。
	CmkKeyId *string `json:"cmk_key_id,omitempty"`

	// **参数解释：** 记录id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 加密主密钥key的状态。 **取值范围：** 1表示待激活状态,2 表示启用状态,3 表示禁用状态,4 表示计划删除状态,5 表示等待导入状态。
	KeyState *UpdateTenantRepoEncryptionSettingResponseKeyState `json:"key_state,omitempty"`

	// **参数解释：** 当前region 。 **取值范围：** 字符串长度不少于1，不超过1000。
	Region *string `json:"region,omitempty"`

	// **参数解释：** region类型。 **取值范围：** 字符串长度不少于1，不超过1000。
	RegionType     *string `json:"region_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTenantRepoEncryptionSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantRepoEncryptionSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantRepoEncryptionSettingResponse", string(data)}, " ")
}

type UpdateTenantRepoEncryptionSettingResponseKeyState struct {
	value string
}

type UpdateTenantRepoEncryptionSettingResponseKeyStateEnum struct {
	E_1 UpdateTenantRepoEncryptionSettingResponseKeyState
	E_2 UpdateTenantRepoEncryptionSettingResponseKeyState
	E_3 UpdateTenantRepoEncryptionSettingResponseKeyState
	E_4 UpdateTenantRepoEncryptionSettingResponseKeyState
	E_5 UpdateTenantRepoEncryptionSettingResponseKeyState
}

func GetUpdateTenantRepoEncryptionSettingResponseKeyStateEnum() UpdateTenantRepoEncryptionSettingResponseKeyStateEnum {
	return UpdateTenantRepoEncryptionSettingResponseKeyStateEnum{
		E_1: UpdateTenantRepoEncryptionSettingResponseKeyState{
			value: "1",
		},
		E_2: UpdateTenantRepoEncryptionSettingResponseKeyState{
			value: "2",
		},
		E_3: UpdateTenantRepoEncryptionSettingResponseKeyState{
			value: "3",
		},
		E_4: UpdateTenantRepoEncryptionSettingResponseKeyState{
			value: "4",
		},
		E_5: UpdateTenantRepoEncryptionSettingResponseKeyState{
			value: "5",
		},
	}
}

func (c UpdateTenantRepoEncryptionSettingResponseKeyState) Value() string {
	return c.value
}

func (c UpdateTenantRepoEncryptionSettingResponseKeyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTenantRepoEncryptionSettingResponseKeyState) UnmarshalJSON(b []byte) error {
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
