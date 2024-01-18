package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateVgwIkePolicy struct {

	// 加密算法
	AuthenticationAlgorithm *UpdateVgwIkePolicyAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`

	// 加密算法
	EncryptionAlgorithm *UpdateVgwIkePolicyEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	// DH密钥组
	DhGroup *UpdateVgwIkePolicyDhGroup `json:"dh_group,omitempty"`

	// 表示SA的生存周期，当该生存周期超时后IKE SA将自动更新
	LifetimeSeconds *int32 `json:"lifetime_seconds,omitempty"`
}

func (o UpdateVgwIkePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwIkePolicy struct{}"
	}

	return strings.Join([]string{"UpdateVgwIkePolicy", string(data)}, " ")
}

type UpdateVgwIkePolicyAuthenticationAlgorithm struct {
	value string
}

type UpdateVgwIkePolicyAuthenticationAlgorithmEnum struct {
	SHA2_256 UpdateVgwIkePolicyAuthenticationAlgorithm
	SHA2_384 UpdateVgwIkePolicyAuthenticationAlgorithm
	SHA2_512 UpdateVgwIkePolicyAuthenticationAlgorithm
}

func GetUpdateVgwIkePolicyAuthenticationAlgorithmEnum() UpdateVgwIkePolicyAuthenticationAlgorithmEnum {
	return UpdateVgwIkePolicyAuthenticationAlgorithmEnum{
		SHA2_256: UpdateVgwIkePolicyAuthenticationAlgorithm{
			value: "sha2-256",
		},
		SHA2_384: UpdateVgwIkePolicyAuthenticationAlgorithm{
			value: "sha2-384",
		},
		SHA2_512: UpdateVgwIkePolicyAuthenticationAlgorithm{
			value: "sha2-512",
		},
	}
}

func (c UpdateVgwIkePolicyAuthenticationAlgorithm) Value() string {
	return c.value
}

func (c UpdateVgwIkePolicyAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVgwIkePolicyAuthenticationAlgorithm) UnmarshalJSON(b []byte) error {
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

type UpdateVgwIkePolicyEncryptionAlgorithm struct {
	value string
}

type UpdateVgwIkePolicyEncryptionAlgorithmEnum struct {
	AES_128        UpdateVgwIkePolicyEncryptionAlgorithm
	AES_192        UpdateVgwIkePolicyEncryptionAlgorithm
	AES_256        UpdateVgwIkePolicyEncryptionAlgorithm
	AES_128_GCM_16 UpdateVgwIkePolicyEncryptionAlgorithm
	AES_256_GCM_16 UpdateVgwIkePolicyEncryptionAlgorithm
}

func GetUpdateVgwIkePolicyEncryptionAlgorithmEnum() UpdateVgwIkePolicyEncryptionAlgorithmEnum {
	return UpdateVgwIkePolicyEncryptionAlgorithmEnum{
		AES_128: UpdateVgwIkePolicyEncryptionAlgorithm{
			value: "aes-128",
		},
		AES_192: UpdateVgwIkePolicyEncryptionAlgorithm{
			value: "aes-192",
		},
		AES_256: UpdateVgwIkePolicyEncryptionAlgorithm{
			value: "aes-256",
		},
		AES_128_GCM_16: UpdateVgwIkePolicyEncryptionAlgorithm{
			value: "aes-128-gcm-16",
		},
		AES_256_GCM_16: UpdateVgwIkePolicyEncryptionAlgorithm{
			value: "aes-256-gcm-16",
		},
	}
}

func (c UpdateVgwIkePolicyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c UpdateVgwIkePolicyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVgwIkePolicyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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

type UpdateVgwIkePolicyDhGroup struct {
	value string
}

type UpdateVgwIkePolicyDhGroupEnum struct {
	GROUP14 UpdateVgwIkePolicyDhGroup
	GROUP15 UpdateVgwIkePolicyDhGroup
	GROUP16 UpdateVgwIkePolicyDhGroup
	GROUP19 UpdateVgwIkePolicyDhGroup
	GROUP20 UpdateVgwIkePolicyDhGroup
	GROUP21 UpdateVgwIkePolicyDhGroup
}

func GetUpdateVgwIkePolicyDhGroupEnum() UpdateVgwIkePolicyDhGroupEnum {
	return UpdateVgwIkePolicyDhGroupEnum{
		GROUP14: UpdateVgwIkePolicyDhGroup{
			value: "group14",
		},
		GROUP15: UpdateVgwIkePolicyDhGroup{
			value: "group15",
		},
		GROUP16: UpdateVgwIkePolicyDhGroup{
			value: "group16",
		},
		GROUP19: UpdateVgwIkePolicyDhGroup{
			value: "group19",
		},
		GROUP20: UpdateVgwIkePolicyDhGroup{
			value: "group20",
		},
		GROUP21: UpdateVgwIkePolicyDhGroup{
			value: "group21",
		},
	}
}

func (c UpdateVgwIkePolicyDhGroup) Value() string {
	return c.value
}

func (c UpdateVgwIkePolicyDhGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVgwIkePolicyDhGroup) UnmarshalJSON(b []byte) error {
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
