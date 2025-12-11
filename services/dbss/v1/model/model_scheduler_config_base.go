package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SchedulerConfigBase 调度配置
type SchedulerConfigBase struct {

	// 数据库ID
	DbIds *string `json:"db_ids,omitempty"`

	// 文件类型
	Format *SchedulerConfigBaseFormat `json:"format,omitempty"`

	// 周期 - AUDIT_REPORT_DAY: 天 - AUDIT_REPORT_WEEK：周 - AUDIT_REPORT_MONTH：月 - AUDIT_REPORT_YEAR：年 - AUDIT_REPORT_REAL_TIME：实时
	Frequency *SchedulerConfigBaseFrequency `json:"frequency,omitempty"`

	// 报表模板ID
	Id *string `json:"id,omitempty"`

	// 调度方式
	Mode *SchedulerConfigBaseMode `json:"mode,omitempty"`

	// 是否通知  - OFF：不通知  - ON：通知
	Notice *SchedulerConfigBaseNotice `json:"notice,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 模板状态 - ON: 开启 - OFF：关闭
	Status *SchedulerConfigBaseStatus `json:"status,omitempty"`

	// 主题URN
	TopicUrn *string `json:"topic_urn,omitempty"`
}

func (o SchedulerConfigBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchedulerConfigBase struct{}"
	}

	return strings.Join([]string{"SchedulerConfigBase", string(data)}, " ")
}

type SchedulerConfigBaseFormat struct {
	value string
}

type SchedulerConfigBaseFormatEnum struct {
	PDF SchedulerConfigBaseFormat
	ZIP SchedulerConfigBaseFormat
}

func GetSchedulerConfigBaseFormatEnum() SchedulerConfigBaseFormatEnum {
	return SchedulerConfigBaseFormatEnum{
		PDF: SchedulerConfigBaseFormat{
			value: "PDF",
		},
		ZIP: SchedulerConfigBaseFormat{
			value: "ZIP",
		},
	}
}

func (c SchedulerConfigBaseFormat) Value() string {
	return c.value
}

func (c SchedulerConfigBaseFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SchedulerConfigBaseFormat) UnmarshalJSON(b []byte) error {
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

type SchedulerConfigBaseFrequency struct {
	value string
}

type SchedulerConfigBaseFrequencyEnum struct {
	AUDIT_REPORT_DAY       SchedulerConfigBaseFrequency
	AUDIT_REPORT_WEEK      SchedulerConfigBaseFrequency
	AUDIT_REPORT_MONTH     SchedulerConfigBaseFrequency
	AUDIT_REPORT_YEAR      SchedulerConfigBaseFrequency
	AUDIT_REPORT_REAL_TIME SchedulerConfigBaseFrequency
}

func GetSchedulerConfigBaseFrequencyEnum() SchedulerConfigBaseFrequencyEnum {
	return SchedulerConfigBaseFrequencyEnum{
		AUDIT_REPORT_DAY: SchedulerConfigBaseFrequency{
			value: "AUDIT_REPORT_DAY",
		},
		AUDIT_REPORT_WEEK: SchedulerConfigBaseFrequency{
			value: "AUDIT_REPORT_WEEK",
		},
		AUDIT_REPORT_MONTH: SchedulerConfigBaseFrequency{
			value: "AUDIT_REPORT_MONTH",
		},
		AUDIT_REPORT_YEAR: SchedulerConfigBaseFrequency{
			value: "AUDIT_REPORT_YEAR",
		},
		AUDIT_REPORT_REAL_TIME: SchedulerConfigBaseFrequency{
			value: "AUDIT_REPORT_REAL_TIME",
		},
	}
}

func (c SchedulerConfigBaseFrequency) Value() string {
	return c.value
}

func (c SchedulerConfigBaseFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SchedulerConfigBaseFrequency) UnmarshalJSON(b []byte) error {
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

type SchedulerConfigBaseMode struct {
	value string
}

type SchedulerConfigBaseModeEnum struct {
	ONCE  SchedulerConfigBaseMode
	CYCLE SchedulerConfigBaseMode
}

func GetSchedulerConfigBaseModeEnum() SchedulerConfigBaseModeEnum {
	return SchedulerConfigBaseModeEnum{
		ONCE: SchedulerConfigBaseMode{
			value: "ONCE",
		},
		CYCLE: SchedulerConfigBaseMode{
			value: "CYCLE",
		},
	}
}

func (c SchedulerConfigBaseMode) Value() string {
	return c.value
}

func (c SchedulerConfigBaseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SchedulerConfigBaseMode) UnmarshalJSON(b []byte) error {
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

type SchedulerConfigBaseNotice struct {
	value string
}

type SchedulerConfigBaseNoticeEnum struct {
	OFF SchedulerConfigBaseNotice
	ON  SchedulerConfigBaseNotice
}

func GetSchedulerConfigBaseNoticeEnum() SchedulerConfigBaseNoticeEnum {
	return SchedulerConfigBaseNoticeEnum{
		OFF: SchedulerConfigBaseNotice{
			value: "OFF",
		},
		ON: SchedulerConfigBaseNotice{
			value: "ON",
		},
	}
}

func (c SchedulerConfigBaseNotice) Value() string {
	return c.value
}

func (c SchedulerConfigBaseNotice) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SchedulerConfigBaseNotice) UnmarshalJSON(b []byte) error {
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

type SchedulerConfigBaseStatus struct {
	value string
}

type SchedulerConfigBaseStatusEnum struct {
	OFF SchedulerConfigBaseStatus
	ON  SchedulerConfigBaseStatus
}

func GetSchedulerConfigBaseStatusEnum() SchedulerConfigBaseStatusEnum {
	return SchedulerConfigBaseStatusEnum{
		OFF: SchedulerConfigBaseStatus{
			value: "OFF",
		},
		ON: SchedulerConfigBaseStatus{
			value: "ON",
		},
	}
}

func (c SchedulerConfigBaseStatus) Value() string {
	return c.value
}

func (c SchedulerConfigBaseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SchedulerConfigBaseStatus) UnmarshalJSON(b []byte) error {
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
