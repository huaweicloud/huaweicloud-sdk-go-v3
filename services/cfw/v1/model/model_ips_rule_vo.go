package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IpsRuleVo struct {
	AffectedApplication *string `json:"affected_application,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	DefaultStatus *IpsRuleVoDefaultStatus `json:"default_status,omitempty"`

	IpsCve *string `json:"ips_cve,omitempty"`

	IpsGroup *IpsRuleVoIpsGroup `json:"ips_group,omitempty"`

	IpsId *string `json:"ips_id,omitempty"`

	IpsLevel *IpsRuleVoIpsLevel `json:"ips_level,omitempty"`

	IpsName *string `json:"ips_name,omitempty"`

	IpsRulesType *string `json:"ips_rules_type,omitempty"`

	IpsStatus *IpsRuleVoIpsStatus `json:"ips_status,omitempty"`
}

func (o IpsRuleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsRuleVo struct{}"
	}

	return strings.Join([]string{"IpsRuleVo", string(data)}, " ")
}

type IpsRuleVoDefaultStatus struct {
	value string
}

type IpsRuleVoDefaultStatusEnum struct {
	OBSERVE     IpsRuleVoDefaultStatus
	ENABLE      IpsRuleVoDefaultStatus
	CLOSE       IpsRuleVoDefaultStatus
	DEFAULT     IpsRuleVoDefaultStatus
	ALL_DEFAULT IpsRuleVoDefaultStatus
}

func GetIpsRuleVoDefaultStatusEnum() IpsRuleVoDefaultStatusEnum {
	return IpsRuleVoDefaultStatusEnum{
		OBSERVE: IpsRuleVoDefaultStatus{
			value: "OBSERVE",
		},
		ENABLE: IpsRuleVoDefaultStatus{
			value: "ENABLE",
		},
		CLOSE: IpsRuleVoDefaultStatus{
			value: "CLOSE",
		},
		DEFAULT: IpsRuleVoDefaultStatus{
			value: "DEFAULT",
		},
		ALL_DEFAULT: IpsRuleVoDefaultStatus{
			value: "ALL_DEFAULT",
		},
	}
}

func (c IpsRuleVoDefaultStatus) Value() string {
	return c.value
}

func (c IpsRuleVoDefaultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsRuleVoDefaultStatus) UnmarshalJSON(b []byte) error {
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

type IpsRuleVoIpsGroup struct {
	value string
}

type IpsRuleVoIpsGroupEnum struct {
	OBSERVE  IpsRuleVoIpsGroup
	STRICTLY IpsRuleVoIpsGroup
	MEDIUM   IpsRuleVoIpsGroup
	LOW      IpsRuleVoIpsGroup
}

func GetIpsRuleVoIpsGroupEnum() IpsRuleVoIpsGroupEnum {
	return IpsRuleVoIpsGroupEnum{
		OBSERVE: IpsRuleVoIpsGroup{
			value: "OBSERVE",
		},
		STRICTLY: IpsRuleVoIpsGroup{
			value: "STRICTLY",
		},
		MEDIUM: IpsRuleVoIpsGroup{
			value: "MEDIUM",
		},
		LOW: IpsRuleVoIpsGroup{
			value: "LOW",
		},
	}
}

func (c IpsRuleVoIpsGroup) Value() string {
	return c.value
}

func (c IpsRuleVoIpsGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsRuleVoIpsGroup) UnmarshalJSON(b []byte) error {
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

type IpsRuleVoIpsLevel struct {
	value string
}

type IpsRuleVoIpsLevelEnum struct {
	FATAL    IpsRuleVoIpsLevel
	HIGH     IpsRuleVoIpsLevel
	MEDIUM   IpsRuleVoIpsLevel
	LOW      IpsRuleVoIpsLevel
	UNKNOWNS IpsRuleVoIpsLevel
}

func GetIpsRuleVoIpsLevelEnum() IpsRuleVoIpsLevelEnum {
	return IpsRuleVoIpsLevelEnum{
		FATAL: IpsRuleVoIpsLevel{
			value: "FATAL",
		},
		HIGH: IpsRuleVoIpsLevel{
			value: "HIGH",
		},
		MEDIUM: IpsRuleVoIpsLevel{
			value: "MEDIUM",
		},
		LOW: IpsRuleVoIpsLevel{
			value: "LOW",
		},
		UNKNOWNS: IpsRuleVoIpsLevel{
			value: "UNKNOWNS",
		},
	}
}

func (c IpsRuleVoIpsLevel) Value() string {
	return c.value
}

func (c IpsRuleVoIpsLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsRuleVoIpsLevel) UnmarshalJSON(b []byte) error {
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

type IpsRuleVoIpsStatus struct {
	value string
}

type IpsRuleVoIpsStatusEnum struct {
	OBSERVE     IpsRuleVoIpsStatus
	ENABLE      IpsRuleVoIpsStatus
	CLOSE       IpsRuleVoIpsStatus
	DEFAULT     IpsRuleVoIpsStatus
	ALL_DEFAULT IpsRuleVoIpsStatus
}

func GetIpsRuleVoIpsStatusEnum() IpsRuleVoIpsStatusEnum {
	return IpsRuleVoIpsStatusEnum{
		OBSERVE: IpsRuleVoIpsStatus{
			value: "OBSERVE",
		},
		ENABLE: IpsRuleVoIpsStatus{
			value: "ENABLE",
		},
		CLOSE: IpsRuleVoIpsStatus{
			value: "CLOSE",
		},
		DEFAULT: IpsRuleVoIpsStatus{
			value: "DEFAULT",
		},
		ALL_DEFAULT: IpsRuleVoIpsStatus{
			value: "ALL_DEFAULT",
		},
	}
}

func (c IpsRuleVoIpsStatus) Value() string {
	return c.value
}

func (c IpsRuleVoIpsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsRuleVoIpsStatus) UnmarshalJSON(b []byte) error {
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
