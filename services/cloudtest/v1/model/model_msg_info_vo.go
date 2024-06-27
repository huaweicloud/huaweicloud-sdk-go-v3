package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type MsgInfoVo struct {

	// 告警渠道
	AlertChannels *string `json:"alert_channels,omitempty"`

	// 告警分组
	AlertGroups *string `json:"alert_groups,omitempty"`

	// 告警级别
	AlertLevel *int32 `json:"alert_level,omitempty"`

	// 总告警次数
	AlertNum *int32 `json:"alert_num,omitempty"`

	// 告警时间
	AlertTime *sdktime.SdkTime `json:"alert_time,omitempty"`

	// 告警类型
	AlertType *MsgInfoVoAlertType `json:"alert_type,omitempty"`

	// 告警内容
	Content *string `json:"content,omitempty"`

	// MsgInfo的id
	Id *string `json:"id,omitempty"`

	// 链接
	Link *string `json:"link,omitempty"`

	// 执行机区域
	LocationNames *string `json:"location_names,omitempty"`

	// 发送告警类型
	SendAlertType *MsgInfoVoSendAlertType `json:"send_alert_type,omitempty"`

	// 子任务用例id
	SubTaskCaseId *string `json:"sub_task_case_id,omitempty"`

	// 子任务用例开始时间
	SubTaskCaseStartTime *sdktime.SdkTime `json:"sub_task_case_start_time,omitempty"`

	// 子任务id
	SubTaskId *string `json:"sub_task_id,omitempty"`

	// 子任务开始时间
	SubTaskStartTime *sdktime.SdkTime `json:"sub_task_start_time,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务名
	TaskName *string `json:"task_name,omitempty"`

	// 用例id
	TestCaseId *string `json:"test_case_id,omitempty"`

	// 用例名称
	TestCaseName *string `json:"test_case_name,omitempty"`

	// 测试套类型
	TestSuiteType *int32 `json:"test_suite_type,omitempty"`
}

func (o MsgInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MsgInfoVo struct{}"
	}

	return strings.Join([]string{"MsgInfoVo", string(data)}, " ")
}

type MsgInfoVoAlertType struct {
	value string
}

type MsgInfoVoAlertTypeEnum struct {
	E_1  MsgInfoVoAlertType
	E_10 MsgInfoVoAlertType
	E_11 MsgInfoVoAlertType
	E_12 MsgInfoVoAlertType
	E_14 MsgInfoVoAlertType
	E_16 MsgInfoVoAlertType
	E_2  MsgInfoVoAlertType
	E_3  MsgInfoVoAlertType
	E_4  MsgInfoVoAlertType
	E_5  MsgInfoVoAlertType
	E_6  MsgInfoVoAlertType
	E_7  MsgInfoVoAlertType
	E_9  MsgInfoVoAlertType
}

func GetMsgInfoVoAlertTypeEnum() MsgInfoVoAlertTypeEnum {
	return MsgInfoVoAlertTypeEnum{
		E_1: MsgInfoVoAlertType{
			value: "1",
		},
		E_10: MsgInfoVoAlertType{
			value: "10",
		},
		E_11: MsgInfoVoAlertType{
			value: "11",
		},
		E_12: MsgInfoVoAlertType{
			value: "12",
		},
		E_14: MsgInfoVoAlertType{
			value: "14",
		},
		E_16: MsgInfoVoAlertType{
			value: "16",
		},
		E_2: MsgInfoVoAlertType{
			value: "2",
		},
		E_3: MsgInfoVoAlertType{
			value: "3",
		},
		E_4: MsgInfoVoAlertType{
			value: "4",
		},
		E_5: MsgInfoVoAlertType{
			value: "5",
		},
		E_6: MsgInfoVoAlertType{
			value: "6",
		},
		E_7: MsgInfoVoAlertType{
			value: "7",
		},
		E_9: MsgInfoVoAlertType{
			value: "9",
		},
	}
}

func (c MsgInfoVoAlertType) Value() string {
	return c.value
}

func (c MsgInfoVoAlertType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MsgInfoVoAlertType) UnmarshalJSON(b []byte) error {
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

type MsgInfoVoSendAlertType struct {
	value string
}

type MsgInfoVoSendAlertTypeEnum struct {
	LEVEL_ALERT      MsgInfoVoSendAlertType
	OPERATION_NOTICE MsgInfoVoSendAlertType
	RECOVER_ALERT    MsgInfoVoSendAlertType
	RESTRAINING      MsgInfoVoSendAlertType
	RESTRAIN_ALERT   MsgInfoVoSendAlertType
}

func GetMsgInfoVoSendAlertTypeEnum() MsgInfoVoSendAlertTypeEnum {
	return MsgInfoVoSendAlertTypeEnum{
		LEVEL_ALERT: MsgInfoVoSendAlertType{
			value: "LEVEL_ALERT",
		},
		OPERATION_NOTICE: MsgInfoVoSendAlertType{
			value: "OPERATION_NOTICE",
		},
		RECOVER_ALERT: MsgInfoVoSendAlertType{
			value: "RECOVER_ALERT",
		},
		RESTRAINING: MsgInfoVoSendAlertType{
			value: "RESTRAINING",
		},
		RESTRAIN_ALERT: MsgInfoVoSendAlertType{
			value: "RESTRAIN_ALERT",
		},
	}
}

func (c MsgInfoVoSendAlertType) Value() string {
	return c.value
}

func (c MsgInfoVoSendAlertType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MsgInfoVoSendAlertType) UnmarshalJSON(b []byte) error {
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
