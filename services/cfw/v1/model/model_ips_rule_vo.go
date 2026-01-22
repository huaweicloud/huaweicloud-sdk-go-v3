package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IpsRuleVo struct {

	// 受影响对象，可包含如下：Others、Sun、Apache、IBM、VMware、WordPress、Adobe、Oracle、Google Chrome等
	AffectedApplication *string `json:"affected_application,omitempty"`

	// ips规则创建的年份
	CreateTime *string `json:"create_time,omitempty"`

	// 默认状态
	DefaultStatus *IpsRuleVoDefaultStatus `json:"default_status,omitempty"`

	// cve id
	IpsCve *string `json:"ips_cve,omitempty"`

	// ips组，使用ips规则拦截模式区分，包含，0：观察模式，1：严格模式，2：中等模式，3：宽松模式
	IpsGroup *IpsRuleVoIpsGroup `json:"ips_group,omitempty"`

	// ips规则id
	IpsId *string `json:"ips_id,omitempty"`

	// ips严重等级，  ips严重等级，包含CRITICAL、HIGH、MEDIUM、LOW
	IpsLevel *IpsRuleVoIpsLevel `json:"ips_level,omitempty"`

	// ips规则名称
	IpsName *string `json:"ips_name,omitempty"`

	// ips规则类型，包括漏洞扫描、黑客工具、特洛伊木马等
	IpsRulesType *string `json:"ips_rules_type,omitempty"`

	// ips规则状态，包含观察：OBSERVE、拦截：ENABLE、禁用：CLOSE、恢复默认：DEFAULT、全局恢复默认：ALL_DEFAULT
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
