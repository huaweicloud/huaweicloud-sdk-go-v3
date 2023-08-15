package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAgencyOption
type CreateAgencyOption struct {

	// 委托名，长度不大于64位。
	Name string `json:"name"`

	// 委托方账号ID。
	DomainId string `json:"domain_id"`

	// 被委托方账号ID。trust_domain_id和trust_domain_name至少填写一个，若都填写优先校验trust_domain_name。
	TrustDomainId *string `json:"trust_domain_id,omitempty"`

	// 被委托方账号名。trust_domain_id和trust_domain_name至少填写一个，若都填写优先校验trust_domain_name。
	TrustDomainName *string `json:"trust_domain_name,omitempty"`

	// 委托描述信息，长度不大于255位。
	Description *string `json:"description,omitempty"`

	// 委托的期限。取值为“FOREVER\"表示委托的期限为永久，取值为\"ONEDAY\"表示委托的期限为一天。不填写该参数则默认为\"null\"也表示委托的期限为永久。
	Duration *CreateAgencyOptionDuration `json:"duration,omitempty"`
}

func (o CreateAgencyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyOption struct{}"
	}

	return strings.Join([]string{"CreateAgencyOption", string(data)}, " ")
}

type CreateAgencyOptionDuration struct {
	value string
}

type CreateAgencyOptionDurationEnum struct {
	FOREVER CreateAgencyOptionDuration
	ONEDAY  CreateAgencyOptionDuration
}

func GetCreateAgencyOptionDurationEnum() CreateAgencyOptionDurationEnum {
	return CreateAgencyOptionDurationEnum{
		FOREVER: CreateAgencyOptionDuration{
			value: "FOREVER",
		},
		ONEDAY: CreateAgencyOptionDuration{
			value: "ONEDAY",
		},
	}
}

func (c CreateAgencyOptionDuration) Value() string {
	return c.value
}

func (c CreateAgencyOptionDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAgencyOptionDuration) UnmarshalJSON(b []byte) error {
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
