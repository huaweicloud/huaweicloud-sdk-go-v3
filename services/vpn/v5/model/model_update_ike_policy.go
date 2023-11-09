package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateIkePolicy struct {

	// IKE协商版本
	IkeVersion *UpdateIkePolicyIkeVersion `json:"ike_version,omitempty"`

	// 协商模式，ike版本为v1时生效
	Phase1NegotiationMode *UpdateIkePolicyPhase1NegotiationMode `json:"phase1_negotiation_mode,omitempty"`

	// 认证算法，SHA1和MD5安全性较低，请慎用
	AuthenticationAlgorithm *UpdateIkePolicyAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`

	// 加密算法，3DES安全性较低，请慎用
	EncryptionAlgorithm *UpdateIkePolicyEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	// DH密钥组
	DhGroup *string `json:"dh_group,omitempty"`

	// 表示SA的生存周期，当该生存周期超时后IKE SA将自动更新
	LifetimeSeconds *int32 `json:"lifetime_seconds,omitempty"`

	// 本端ID类型
	LocalIdType *UpdateIkePolicyLocalIdType `json:"local_id_type,omitempty"`

	// 本端ID
	LocalId *string `json:"local_id,omitempty"`

	// 对端ID类型
	PeerIdType *UpdateIkePolicyPeerIdType `json:"peer_id_type,omitempty"`

	// 对端ID
	PeerId *string `json:"peer_id,omitempty"`

	Dpd *UpdateDpd `json:"dpd,omitempty"`
}

func (o UpdateIkePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIkePolicy struct{}"
	}

	return strings.Join([]string{"UpdateIkePolicy", string(data)}, " ")
}

type UpdateIkePolicyIkeVersion struct {
	value string
}

type UpdateIkePolicyIkeVersionEnum struct {
	V1 UpdateIkePolicyIkeVersion
	V2 UpdateIkePolicyIkeVersion
}

func GetUpdateIkePolicyIkeVersionEnum() UpdateIkePolicyIkeVersionEnum {
	return UpdateIkePolicyIkeVersionEnum{
		V1: UpdateIkePolicyIkeVersion{
			value: "v1",
		},
		V2: UpdateIkePolicyIkeVersion{
			value: "v2",
		},
	}
}

func (c UpdateIkePolicyIkeVersion) Value() string {
	return c.value
}

func (c UpdateIkePolicyIkeVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIkePolicyIkeVersion) UnmarshalJSON(b []byte) error {
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

type UpdateIkePolicyPhase1NegotiationMode struct {
	value string
}

type UpdateIkePolicyPhase1NegotiationModeEnum struct {
	MAIN       UpdateIkePolicyPhase1NegotiationMode
	AGGRESSIVE UpdateIkePolicyPhase1NegotiationMode
}

func GetUpdateIkePolicyPhase1NegotiationModeEnum() UpdateIkePolicyPhase1NegotiationModeEnum {
	return UpdateIkePolicyPhase1NegotiationModeEnum{
		MAIN: UpdateIkePolicyPhase1NegotiationMode{
			value: "main",
		},
		AGGRESSIVE: UpdateIkePolicyPhase1NegotiationMode{
			value: "aggressive",
		},
	}
}

func (c UpdateIkePolicyPhase1NegotiationMode) Value() string {
	return c.value
}

func (c UpdateIkePolicyPhase1NegotiationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIkePolicyPhase1NegotiationMode) UnmarshalJSON(b []byte) error {
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

type UpdateIkePolicyAuthenticationAlgorithm struct {
	value string
}

type UpdateIkePolicyAuthenticationAlgorithmEnum struct {
	SHA1     UpdateIkePolicyAuthenticationAlgorithm
	MD5      UpdateIkePolicyAuthenticationAlgorithm
	SHA2_256 UpdateIkePolicyAuthenticationAlgorithm
	SHA2_384 UpdateIkePolicyAuthenticationAlgorithm
	SHA2_512 UpdateIkePolicyAuthenticationAlgorithm
}

func GetUpdateIkePolicyAuthenticationAlgorithmEnum() UpdateIkePolicyAuthenticationAlgorithmEnum {
	return UpdateIkePolicyAuthenticationAlgorithmEnum{
		SHA1: UpdateIkePolicyAuthenticationAlgorithm{
			value: "sha1",
		},
		MD5: UpdateIkePolicyAuthenticationAlgorithm{
			value: "md5",
		},
		SHA2_256: UpdateIkePolicyAuthenticationAlgorithm{
			value: "sha2-256",
		},
		SHA2_384: UpdateIkePolicyAuthenticationAlgorithm{
			value: "sha2-384",
		},
		SHA2_512: UpdateIkePolicyAuthenticationAlgorithm{
			value: "sha2-512",
		},
	}
}

func (c UpdateIkePolicyAuthenticationAlgorithm) Value() string {
	return c.value
}

func (c UpdateIkePolicyAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIkePolicyAuthenticationAlgorithm) UnmarshalJSON(b []byte) error {
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

type UpdateIkePolicyEncryptionAlgorithm struct {
	value string
}

type UpdateIkePolicyEncryptionAlgorithmEnum struct {
	E_3DES          UpdateIkePolicyEncryptionAlgorithm
	AES_128         UpdateIkePolicyEncryptionAlgorithm
	AES_192         UpdateIkePolicyEncryptionAlgorithm
	AES_256         UpdateIkePolicyEncryptionAlgorithm
	AES_128_GCM_16  UpdateIkePolicyEncryptionAlgorithm
	AES_256_GCM_16  UpdateIkePolicyEncryptionAlgorithm
	AES_128_GCM_128 UpdateIkePolicyEncryptionAlgorithm
	AES_256_GCM_128 UpdateIkePolicyEncryptionAlgorithm
}

func GetUpdateIkePolicyEncryptionAlgorithmEnum() UpdateIkePolicyEncryptionAlgorithmEnum {
	return UpdateIkePolicyEncryptionAlgorithmEnum{
		E_3DES: UpdateIkePolicyEncryptionAlgorithm{
			value: "3des",
		},
		AES_128: UpdateIkePolicyEncryptionAlgorithm{
			value: "aes-128",
		},
		AES_192: UpdateIkePolicyEncryptionAlgorithm{
			value: "aes-192",
		},
		AES_256: UpdateIkePolicyEncryptionAlgorithm{
			value: "aes-256",
		},
		AES_128_GCM_16: UpdateIkePolicyEncryptionAlgorithm{
			value: "aes-128-gcm-16",
		},
		AES_256_GCM_16: UpdateIkePolicyEncryptionAlgorithm{
			value: "aes-256-gcm-16",
		},
		AES_128_GCM_128: UpdateIkePolicyEncryptionAlgorithm{
			value: "aes-128-gcm-128",
		},
		AES_256_GCM_128: UpdateIkePolicyEncryptionAlgorithm{
			value: "aes-256-gcm-128",
		},
	}
}

func (c UpdateIkePolicyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c UpdateIkePolicyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIkePolicyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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

type UpdateIkePolicyLocalIdType struct {
	value string
}

type UpdateIkePolicyLocalIdTypeEnum struct {
	IP   UpdateIkePolicyLocalIdType
	FQDN UpdateIkePolicyLocalIdType
}

func GetUpdateIkePolicyLocalIdTypeEnum() UpdateIkePolicyLocalIdTypeEnum {
	return UpdateIkePolicyLocalIdTypeEnum{
		IP: UpdateIkePolicyLocalIdType{
			value: "ip",
		},
		FQDN: UpdateIkePolicyLocalIdType{
			value: "fqdn",
		},
	}
}

func (c UpdateIkePolicyLocalIdType) Value() string {
	return c.value
}

func (c UpdateIkePolicyLocalIdType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIkePolicyLocalIdType) UnmarshalJSON(b []byte) error {
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

type UpdateIkePolicyPeerIdType struct {
	value string
}

type UpdateIkePolicyPeerIdTypeEnum struct {
	IP   UpdateIkePolicyPeerIdType
	FQDN UpdateIkePolicyPeerIdType
}

func GetUpdateIkePolicyPeerIdTypeEnum() UpdateIkePolicyPeerIdTypeEnum {
	return UpdateIkePolicyPeerIdTypeEnum{
		IP: UpdateIkePolicyPeerIdType{
			value: "ip",
		},
		FQDN: UpdateIkePolicyPeerIdType{
			value: "fqdn",
		},
	}
}

func (c UpdateIkePolicyPeerIdType) Value() string {
	return c.value
}

func (c UpdateIkePolicyPeerIdType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIkePolicyPeerIdType) UnmarshalJSON(b []byte) error {
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
