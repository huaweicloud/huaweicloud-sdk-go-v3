package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowDetailsOfDomainNameCertificateV2Response Response Object
type ShowDetailsOfDomainNameCertificateV2Response struct {

	// 证书ID
	Id *string `json:"id,omitempty"`

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 证书类型  - global：全局证书 - instance：实例证书
	Type *ShowDetailsOfDomainNameCertificateV2ResponseType `json:"type,omitempty"`

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

	// 证书算法类型： - RSA - ECC - SM2  [暂不支持](tag:hws;hws_hk;g42;Site)
	AlgorithmType *ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType `json:"algorithm_type,omitempty"`

	// 是否存在信任的根证书CA。当绑定证书存在trusted_root_ca时为true。  [暂不支持](tag:fcs;hcs;g42;Site)
	IsHasTrustedRootCa *bool `json:"is_has_trusted_root_ca,omitempty"`

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

func (o ShowDetailsOfDomainNameCertificateV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfDomainNameCertificateV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfDomainNameCertificateV2Response", string(data)}, " ")
}

type ShowDetailsOfDomainNameCertificateV2ResponseType struct {
	value string
}

type ShowDetailsOfDomainNameCertificateV2ResponseTypeEnum struct {
	GLOBAL   ShowDetailsOfDomainNameCertificateV2ResponseType
	INSTANCE ShowDetailsOfDomainNameCertificateV2ResponseType
}

func GetShowDetailsOfDomainNameCertificateV2ResponseTypeEnum() ShowDetailsOfDomainNameCertificateV2ResponseTypeEnum {
	return ShowDetailsOfDomainNameCertificateV2ResponseTypeEnum{
		GLOBAL: ShowDetailsOfDomainNameCertificateV2ResponseType{
			value: "global",
		},
		INSTANCE: ShowDetailsOfDomainNameCertificateV2ResponseType{
			value: "instance",
		},
	}
}

func (c ShowDetailsOfDomainNameCertificateV2ResponseType) Value() string {
	return c.value
}

func (c ShowDetailsOfDomainNameCertificateV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfDomainNameCertificateV2ResponseType) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType struct {
	value string
}

type ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmTypeEnum struct {
	RSA ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType
	ECC ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType
	SM2 ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType
}

func GetShowDetailsOfDomainNameCertificateV2ResponseAlgorithmTypeEnum() ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmTypeEnum {
	return ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmTypeEnum{
		RSA: ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType{
			value: "RSA",
		},
		ECC: ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType{
			value: "ECC",
		},
		SM2: ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType{
			value: "SM2",
		},
	}
}

func (c ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType) Value() string {
	return c.value
}

func (c ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfDomainNameCertificateV2ResponseAlgorithmType) UnmarshalJSON(b []byte) error {
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
