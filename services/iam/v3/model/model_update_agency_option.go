package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type UpdateAgencyOption struct {
	// 被委托方账号ID。如果trust_domain_id和trust_domain_name都填写，则优先校验trust_domain_name。四个参数至少填写一个。

	TrustDomainId *string `json:"trust_domain_id,omitempty"`
	// 被委托方账号名。如果trust_domain_id和trust_domain_name都填写，则优先校验trust_domain_name。四个参数至少填写一个。

	TrustDomainName *string `json:"trust_domain_name,omitempty"`
	// 委托描述信息，长度不大于255位。四个参数至少填写一个。

	Description *string `json:"description,omitempty"`
	// 委托的期限。取值为“FOREVER\"表示委托的期限为永久，取值为\"ONEDAY\"表示委托的期限为一天。四个参数至少填写一个。

	Duration *UpdateAgencyOptionDuration `json:"duration,omitempty"`
}

func (o UpdateAgencyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyOption struct{}"
	}

	return strings.Join([]string{"UpdateAgencyOption", string(data)}, " ")
}

type UpdateAgencyOptionDuration struct {
	value string
}

type UpdateAgencyOptionDurationEnum struct {
	FOREVER UpdateAgencyOptionDuration
	ONEDAY  UpdateAgencyOptionDuration
}

func GetUpdateAgencyOptionDurationEnum() UpdateAgencyOptionDurationEnum {
	return UpdateAgencyOptionDurationEnum{
		FOREVER: UpdateAgencyOptionDuration{
			value: "FOREVER",
		},
		ONEDAY: UpdateAgencyOptionDuration{
			value: "ONEDAY",
		},
	}
}

func (c UpdateAgencyOptionDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAgencyOptionDuration) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
