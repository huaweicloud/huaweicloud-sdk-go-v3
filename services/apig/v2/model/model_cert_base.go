package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type CertBase struct {

	// 证书ID
	Id *string `json:"id,omitempty"`

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 证书类型  - global：全局证书 - instance：实例证书
	Type *CertBaseType `json:"type,omitempty"`

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
}

func (o CertBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertBase struct{}"
	}

	return strings.Join([]string{"CertBase", string(data)}, " ")
}

type CertBaseType struct {
	value string
}

type CertBaseTypeEnum struct {
	GLOBAL   CertBaseType
	INSTANCE CertBaseType
}

func GetCertBaseTypeEnum() CertBaseTypeEnum {
	return CertBaseTypeEnum{
		GLOBAL: CertBaseType{
			value: "global",
		},
		INSTANCE: CertBaseType{
			value: "instance",
		},
	}
}

func (c CertBaseType) Value() string {
	return c.value
}

func (c CertBaseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertBaseType) UnmarshalJSON(b []byte) error {
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
