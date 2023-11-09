package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IkePolicy struct {

	// IKE协商版本
	IkeVersion *IkePolicyIkeVersion `json:"ike_version,omitempty"`

	// 协商模式，ike版本为v1时生效
	Phase1NegotiationMode *IkePolicyPhase1NegotiationMode `json:"phase1_negotiation_mode,omitempty"`

	// 认证算法，SHA1和MD5安全性较低，请慎用
	AuthenticationAlgorithm *IkePolicyAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`

	// 加密算法，3DES安全性较低，请慎用
	EncryptionAlgorithm *IkePolicyEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	// DH密钥组
	DhGroup *string `json:"dh_group,omitempty"`

	// ike协商时的认证方法
	AuthenticationMethod *IkePolicyAuthenticationMethod `json:"authentication_method,omitempty"`

	// 表示SA的生存周期，当该生存周期超时后IKE SA将自动更新
	LifetimeSeconds *int32 `json:"lifetime_seconds,omitempty"`

	// 本端ID类型
	LocalIdType *IkePolicyLocalIdType `json:"local_id_type,omitempty"`

	// 本端ID
	LocalId *string `json:"local_id,omitempty"`

	// 对端ID类型
	PeerIdType *IkePolicyPeerIdType `json:"peer_id_type,omitempty"`

	// 对端ID
	PeerId *string `json:"peer_id,omitempty"`

	Dpd *Dpd `json:"dpd,omitempty"`
}

func (o IkePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IkePolicy struct{}"
	}

	return strings.Join([]string{"IkePolicy", string(data)}, " ")
}

type IkePolicyIkeVersion struct {
	value string
}

type IkePolicyIkeVersionEnum struct {
	V1 IkePolicyIkeVersion
	V2 IkePolicyIkeVersion
}

func GetIkePolicyIkeVersionEnum() IkePolicyIkeVersionEnum {
	return IkePolicyIkeVersionEnum{
		V1: IkePolicyIkeVersion{
			value: "v1",
		},
		V2: IkePolicyIkeVersion{
			value: "v2",
		},
	}
}

func (c IkePolicyIkeVersion) Value() string {
	return c.value
}

func (c IkePolicyIkeVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IkePolicyIkeVersion) UnmarshalJSON(b []byte) error {
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

type IkePolicyPhase1NegotiationMode struct {
	value string
}

type IkePolicyPhase1NegotiationModeEnum struct {
	MAIN       IkePolicyPhase1NegotiationMode
	AGGRESSIVE IkePolicyPhase1NegotiationMode
}

func GetIkePolicyPhase1NegotiationModeEnum() IkePolicyPhase1NegotiationModeEnum {
	return IkePolicyPhase1NegotiationModeEnum{
		MAIN: IkePolicyPhase1NegotiationMode{
			value: "main",
		},
		AGGRESSIVE: IkePolicyPhase1NegotiationMode{
			value: "aggressive",
		},
	}
}

func (c IkePolicyPhase1NegotiationMode) Value() string {
	return c.value
}

func (c IkePolicyPhase1NegotiationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IkePolicyPhase1NegotiationMode) UnmarshalJSON(b []byte) error {
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

type IkePolicyAuthenticationAlgorithm struct {
	value string
}

type IkePolicyAuthenticationAlgorithmEnum struct {
	SHA1     IkePolicyAuthenticationAlgorithm
	MD5      IkePolicyAuthenticationAlgorithm
	SHA2_256 IkePolicyAuthenticationAlgorithm
	SHA2_384 IkePolicyAuthenticationAlgorithm
	SHA2_512 IkePolicyAuthenticationAlgorithm
}

func GetIkePolicyAuthenticationAlgorithmEnum() IkePolicyAuthenticationAlgorithmEnum {
	return IkePolicyAuthenticationAlgorithmEnum{
		SHA1: IkePolicyAuthenticationAlgorithm{
			value: "sha1",
		},
		MD5: IkePolicyAuthenticationAlgorithm{
			value: "md5",
		},
		SHA2_256: IkePolicyAuthenticationAlgorithm{
			value: "sha2-256",
		},
		SHA2_384: IkePolicyAuthenticationAlgorithm{
			value: "sha2-384",
		},
		SHA2_512: IkePolicyAuthenticationAlgorithm{
			value: "sha2-512",
		},
	}
}

func (c IkePolicyAuthenticationAlgorithm) Value() string {
	return c.value
}

func (c IkePolicyAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IkePolicyAuthenticationAlgorithm) UnmarshalJSON(b []byte) error {
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

type IkePolicyEncryptionAlgorithm struct {
	value string
}

type IkePolicyEncryptionAlgorithmEnum struct {
	E_3DES          IkePolicyEncryptionAlgorithm
	AES_128         IkePolicyEncryptionAlgorithm
	AES_192         IkePolicyEncryptionAlgorithm
	AES_256         IkePolicyEncryptionAlgorithm
	AES_128_GCM_16  IkePolicyEncryptionAlgorithm
	AES_256_GCM_16  IkePolicyEncryptionAlgorithm
	AES_128_GCM_128 IkePolicyEncryptionAlgorithm
	AES_256_GCM_128 IkePolicyEncryptionAlgorithm
}

func GetIkePolicyEncryptionAlgorithmEnum() IkePolicyEncryptionAlgorithmEnum {
	return IkePolicyEncryptionAlgorithmEnum{
		E_3DES: IkePolicyEncryptionAlgorithm{
			value: "3des",
		},
		AES_128: IkePolicyEncryptionAlgorithm{
			value: "aes-128",
		},
		AES_192: IkePolicyEncryptionAlgorithm{
			value: "aes-192",
		},
		AES_256: IkePolicyEncryptionAlgorithm{
			value: "aes-256",
		},
		AES_128_GCM_16: IkePolicyEncryptionAlgorithm{
			value: "aes-128-gcm-16",
		},
		AES_256_GCM_16: IkePolicyEncryptionAlgorithm{
			value: "aes-256-gcm-16",
		},
		AES_128_GCM_128: IkePolicyEncryptionAlgorithm{
			value: "aes-128-gcm-128",
		},
		AES_256_GCM_128: IkePolicyEncryptionAlgorithm{
			value: "aes-256-gcm-128",
		},
	}
}

func (c IkePolicyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c IkePolicyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IkePolicyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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

type IkePolicyAuthenticationMethod struct {
	value string
}

type IkePolicyAuthenticationMethodEnum struct {
	PRE_SHARE           IkePolicyAuthenticationMethod
	DIGITAL_ENVELOPE_V2 IkePolicyAuthenticationMethod
}

func GetIkePolicyAuthenticationMethodEnum() IkePolicyAuthenticationMethodEnum {
	return IkePolicyAuthenticationMethodEnum{
		PRE_SHARE: IkePolicyAuthenticationMethod{
			value: "pre-share",
		},
		DIGITAL_ENVELOPE_V2: IkePolicyAuthenticationMethod{
			value: "digital-envelope-v2",
		},
	}
}

func (c IkePolicyAuthenticationMethod) Value() string {
	return c.value
}

func (c IkePolicyAuthenticationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IkePolicyAuthenticationMethod) UnmarshalJSON(b []byte) error {
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

type IkePolicyLocalIdType struct {
	value string
}

type IkePolicyLocalIdTypeEnum struct {
	IP   IkePolicyLocalIdType
	FQDN IkePolicyLocalIdType
}

func GetIkePolicyLocalIdTypeEnum() IkePolicyLocalIdTypeEnum {
	return IkePolicyLocalIdTypeEnum{
		IP: IkePolicyLocalIdType{
			value: "ip",
		},
		FQDN: IkePolicyLocalIdType{
			value: "fqdn",
		},
	}
}

func (c IkePolicyLocalIdType) Value() string {
	return c.value
}

func (c IkePolicyLocalIdType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IkePolicyLocalIdType) UnmarshalJSON(b []byte) error {
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

type IkePolicyPeerIdType struct {
	value string
}

type IkePolicyPeerIdTypeEnum struct {
	IP   IkePolicyPeerIdType
	FQDN IkePolicyPeerIdType
}

func GetIkePolicyPeerIdTypeEnum() IkePolicyPeerIdTypeEnum {
	return IkePolicyPeerIdTypeEnum{
		IP: IkePolicyPeerIdType{
			value: "ip",
		},
		FQDN: IkePolicyPeerIdType{
			value: "fqdn",
		},
	}
}

func (c IkePolicyPeerIdType) Value() string {
	return c.value
}

func (c IkePolicyPeerIdType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IkePolicyPeerIdType) UnmarshalJSON(b []byte) error {
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
