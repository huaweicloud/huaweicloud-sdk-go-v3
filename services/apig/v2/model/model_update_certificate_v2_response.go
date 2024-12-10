package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateCertificateV2Response Response Object
type UpdateCertificateV2Response struct {

	// 证书ID
	Id *string `json:"id,omitempty"`

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 证书类型  - global：全局证书 - instance：实例证书
	Type *UpdateCertificateV2ResponseType `json:"type,omitempty"`

	// 实例编码  - `type`为`global`时，缺省为common - `type`为`instance`时，为实例编码
	InstanceId *string `json:"instance_id,omitempty"`

	// 租户项目编号
	ProjectId *string `json:"project_id,omitempty"`

	// 域名
	CommonName *string `json:"common_name,omitempty"`

	// san扩展域名
	San *[]string `json:"san,omitempty"`

	// 有效期到
	NotAfter *sdktime.SdkTime `json:"not_after,omitempty"`

	// 签名算法
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 是否存在信任的根证书CA。当绑定证书存在trusted_root_ca时为true。
	IsHasTrustedRootCa *bool `json:"is_has_trusted_root_ca,omitempty"`

	// 证书算法类型： - RSA - ECC - SM2
	AlgorithmType *UpdateCertificateV2ResponseAlgorithmType `json:"algorithm_type,omitempty"`

	// 版本
	Version *int32 `json:"version,omitempty"`

	// 公司、组织
	Organization *[]string `json:"organization,omitempty"`

	// 部门
	OrganizationalUnit *[]string `json:"organizational_unit,omitempty"`

	// 城市
	Locality *[]string `json:"locality,omitempty"`

	// 省份
	State *[]string `json:"state,omitempty"`

	// 国家
	Country *[]string `json:"country,omitempty"`

	// 有效期从
	NotBefore *sdktime.SdkTime `json:"not_before,omitempty"`

	// 序列号
	SerialNumber *string `json:"serial_number,omitempty"`

	// 颁发者
	Issuer         *[]string `json:"issuer,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateCertificateV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateV2Response struct{}"
	}

	return strings.Join([]string{"UpdateCertificateV2Response", string(data)}, " ")
}

type UpdateCertificateV2ResponseType struct {
	value string
}

type UpdateCertificateV2ResponseTypeEnum struct {
	GLOBAL   UpdateCertificateV2ResponseType
	INSTANCE UpdateCertificateV2ResponseType
}

func GetUpdateCertificateV2ResponseTypeEnum() UpdateCertificateV2ResponseTypeEnum {
	return UpdateCertificateV2ResponseTypeEnum{
		GLOBAL: UpdateCertificateV2ResponseType{
			value: "global",
		},
		INSTANCE: UpdateCertificateV2ResponseType{
			value: "instance",
		},
	}
}

func (c UpdateCertificateV2ResponseType) Value() string {
	return c.value
}

func (c UpdateCertificateV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCertificateV2ResponseType) UnmarshalJSON(b []byte) error {
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

type UpdateCertificateV2ResponseAlgorithmType struct {
	value string
}

type UpdateCertificateV2ResponseAlgorithmTypeEnum struct {
	RSA UpdateCertificateV2ResponseAlgorithmType
	ECC UpdateCertificateV2ResponseAlgorithmType
	SM2 UpdateCertificateV2ResponseAlgorithmType
}

func GetUpdateCertificateV2ResponseAlgorithmTypeEnum() UpdateCertificateV2ResponseAlgorithmTypeEnum {
	return UpdateCertificateV2ResponseAlgorithmTypeEnum{
		RSA: UpdateCertificateV2ResponseAlgorithmType{
			value: "RSA",
		},
		ECC: UpdateCertificateV2ResponseAlgorithmType{
			value: "ECC",
		},
		SM2: UpdateCertificateV2ResponseAlgorithmType{
			value: "SM2",
		},
	}
}

func (c UpdateCertificateV2ResponseAlgorithmType) Value() string {
	return c.value
}

func (c UpdateCertificateV2ResponseAlgorithmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCertificateV2ResponseAlgorithmType) UnmarshalJSON(b []byte) error {
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
