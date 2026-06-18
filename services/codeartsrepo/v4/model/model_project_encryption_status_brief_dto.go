package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProjectEncryptionStatusBriefDto struct {

	// **参数解释：** 仓库加密状态ID。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库加密状态 **取值范围：** - encrypting，加密中。 - encrypted，已加密。 - decrypting，解密中。 - decrypted，已解密。 **默认取值：** 不涉及。
	Status *ProjectEncryptionStatusBriefDtoStatus `json:"status,omitempty"`

	// **参数解释：** 最近加密时间。 **取值范围：** 不涉及。
	LastEncryptionAt *string `json:"last_encryption_at,omitempty"`

	// **参数解释：** 最近解密时间。 **取值范围：** 不涉及。
	LastDecryptionAt *string `json:"last_decryption_at,omitempty"`

	// **参数解释：** 是否开启仓库加密。 **约束限制：** 不涉及。 **取值范围：** - true，开启仓库加密。 - false，关闭仓库加密。
	IsConsistent *bool `json:"is_consistent,omitempty"`
}

func (o ProjectEncryptionStatusBriefDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectEncryptionStatusBriefDto struct{}"
	}

	return strings.Join([]string{"ProjectEncryptionStatusBriefDto", string(data)}, " ")
}

type ProjectEncryptionStatusBriefDtoStatus struct {
	value string
}

type ProjectEncryptionStatusBriefDtoStatusEnum struct {
	ENCRYPTING ProjectEncryptionStatusBriefDtoStatus
	ENCRYPTED  ProjectEncryptionStatusBriefDtoStatus
	DECRYPTING ProjectEncryptionStatusBriefDtoStatus
	DECRYPTED  ProjectEncryptionStatusBriefDtoStatus
}

func GetProjectEncryptionStatusBriefDtoStatusEnum() ProjectEncryptionStatusBriefDtoStatusEnum {
	return ProjectEncryptionStatusBriefDtoStatusEnum{
		ENCRYPTING: ProjectEncryptionStatusBriefDtoStatus{
			value: "encrypting",
		},
		ENCRYPTED: ProjectEncryptionStatusBriefDtoStatus{
			value: "encrypted",
		},
		DECRYPTING: ProjectEncryptionStatusBriefDtoStatus{
			value: "decrypting",
		},
		DECRYPTED: ProjectEncryptionStatusBriefDtoStatus{
			value: "decrypted",
		},
	}
}

func (c ProjectEncryptionStatusBriefDtoStatus) Value() string {
	return c.value
}

func (c ProjectEncryptionStatusBriefDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectEncryptionStatusBriefDtoStatus) UnmarshalJSON(b []byte) error {
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
