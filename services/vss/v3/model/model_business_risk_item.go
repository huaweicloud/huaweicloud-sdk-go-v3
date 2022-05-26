package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BusinessRiskItem struct {

	// 业务风险ID
	RiskId *string `json:"risk_id,omitempty"`

	// 有风险的URL
	RiskUrl *string `json:"risk_url,omitempty"`

	// 业务风险类型:   * text - 不合规文字   * image - 不合规图片   * dead_link - 不合规链接（死链）   * dark_link - 不合规链接（暗链）   * business_risk - 业务风险
	RiskType *BusinessRiskItemRiskType `json:"risk_type,omitempty"`

	// 业务风险发现时间
	FindTime *string `json:"find_time,omitempty"`

	// 业务风险内容
	RiskContent *string `json:"risk_content,omitempty"`

	// 漏洞状态:   * repairing - 未修复   * repaired - 已修复   * false_report - 误报，已忽略
	RiskStatus *BusinessRiskItemRiskStatus `json:"risk_status,omitempty"`
}

func (o BusinessRiskItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessRiskItem struct{}"
	}

	return strings.Join([]string{"BusinessRiskItem", string(data)}, " ")
}

type BusinessRiskItemRiskType struct {
	value string
}

type BusinessRiskItemRiskTypeEnum struct {
	TEXT          BusinessRiskItemRiskType
	IMAGE         BusinessRiskItemRiskType
	DEAD_LINK     BusinessRiskItemRiskType
	DARK_LINK     BusinessRiskItemRiskType
	BUSINESS_RISK BusinessRiskItemRiskType
}

func GetBusinessRiskItemRiskTypeEnum() BusinessRiskItemRiskTypeEnum {
	return BusinessRiskItemRiskTypeEnum{
		TEXT: BusinessRiskItemRiskType{
			value: "text",
		},
		IMAGE: BusinessRiskItemRiskType{
			value: "image",
		},
		DEAD_LINK: BusinessRiskItemRiskType{
			value: "dead_link",
		},
		DARK_LINK: BusinessRiskItemRiskType{
			value: "dark_link",
		},
		BUSINESS_RISK: BusinessRiskItemRiskType{
			value: "business_risk",
		},
	}
}

func (c BusinessRiskItemRiskType) Value() string {
	return c.value
}

func (c BusinessRiskItemRiskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BusinessRiskItemRiskType) UnmarshalJSON(b []byte) error {
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

type BusinessRiskItemRiskStatus struct {
	value string
}

type BusinessRiskItemRiskStatusEnum struct {
	REPAIRING    BusinessRiskItemRiskStatus
	REPAIRED     BusinessRiskItemRiskStatus
	FALSE_REPORT BusinessRiskItemRiskStatus
}

func GetBusinessRiskItemRiskStatusEnum() BusinessRiskItemRiskStatusEnum {
	return BusinessRiskItemRiskStatusEnum{
		REPAIRING: BusinessRiskItemRiskStatus{
			value: "repairing",
		},
		REPAIRED: BusinessRiskItemRiskStatus{
			value: "repaired",
		},
		FALSE_REPORT: BusinessRiskItemRiskStatus{
			value: "false_report",
		},
	}
}

func (c BusinessRiskItemRiskStatus) Value() string {
	return c.value
}

func (c BusinessRiskItemRiskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BusinessRiskItemRiskStatus) UnmarshalJSON(b []byte) error {
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
