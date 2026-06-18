package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTenantRepoEncryptionSettingResponse Response Object
type ShowTenantRepoEncryptionSettingResponse struct {

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
	KeyState *ShowTenantRepoEncryptionSettingResponseKeyState `json:"key_state,omitempty"`

	// **参数解释：** 当前region 。 **取值范围：** 字符串长度不少于1，不超过1000。
	Region *string `json:"region,omitempty"`

	// **参数解释：** region类型。 **取值范围：** 字符串长度不少于1，不超过1000。
	RegionType     *string `json:"region_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTenantRepoEncryptionSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantRepoEncryptionSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantRepoEncryptionSettingResponse", string(data)}, " ")
}

type ShowTenantRepoEncryptionSettingResponseKeyState struct {
	value string
}

type ShowTenantRepoEncryptionSettingResponseKeyStateEnum struct {
	E_1 ShowTenantRepoEncryptionSettingResponseKeyState
	E_2 ShowTenantRepoEncryptionSettingResponseKeyState
	E_3 ShowTenantRepoEncryptionSettingResponseKeyState
	E_4 ShowTenantRepoEncryptionSettingResponseKeyState
	E_5 ShowTenantRepoEncryptionSettingResponseKeyState
}

func GetShowTenantRepoEncryptionSettingResponseKeyStateEnum() ShowTenantRepoEncryptionSettingResponseKeyStateEnum {
	return ShowTenantRepoEncryptionSettingResponseKeyStateEnum{
		E_1: ShowTenantRepoEncryptionSettingResponseKeyState{
			value: "1",
		},
		E_2: ShowTenantRepoEncryptionSettingResponseKeyState{
			value: "2",
		},
		E_3: ShowTenantRepoEncryptionSettingResponseKeyState{
			value: "3",
		},
		E_4: ShowTenantRepoEncryptionSettingResponseKeyState{
			value: "4",
		},
		E_5: ShowTenantRepoEncryptionSettingResponseKeyState{
			value: "5",
		},
	}
}

func (c ShowTenantRepoEncryptionSettingResponseKeyState) Value() string {
	return c.value
}

func (c ShowTenantRepoEncryptionSettingResponseKeyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTenantRepoEncryptionSettingResponseKeyState) UnmarshalJSON(b []byte) error {
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
