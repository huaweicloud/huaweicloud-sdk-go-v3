package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MsgInfoQuery struct {

	// 告警类型
	AlertTypes *[]MsgInfoQueryAlertTypes `json:"alert_types,omitempty"`

	// 用例id
	CaseId *string `json:"case_id,omitempty"`

	// 用例名
	CaseName *string `json:"case_name,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 页码
	PageNum *int32 `json:"page_num,omitempty"`

	// 分页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 发送类型
	SendAlertType *MsgInfoQuerySendAlertType `json:"send_alert_type,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务名
	TaskName *string `json:"task_name,omitempty"`
}

func (o MsgInfoQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MsgInfoQuery struct{}"
	}

	return strings.Join([]string{"MsgInfoQuery", string(data)}, " ")
}

type MsgInfoQueryAlertTypes struct {
	value string
}

type MsgInfoQueryAlertTypesEnum struct {
	E_1  MsgInfoQueryAlertTypes
	E_10 MsgInfoQueryAlertTypes
	E_11 MsgInfoQueryAlertTypes
	E_12 MsgInfoQueryAlertTypes
	E_14 MsgInfoQueryAlertTypes
	E_16 MsgInfoQueryAlertTypes
	E_2  MsgInfoQueryAlertTypes
	E_3  MsgInfoQueryAlertTypes
	E_4  MsgInfoQueryAlertTypes
	E_5  MsgInfoQueryAlertTypes
	E_6  MsgInfoQueryAlertTypes
	E_7  MsgInfoQueryAlertTypes
	E_9  MsgInfoQueryAlertTypes
}

func GetMsgInfoQueryAlertTypesEnum() MsgInfoQueryAlertTypesEnum {
	return MsgInfoQueryAlertTypesEnum{
		E_1: MsgInfoQueryAlertTypes{
			value: "1",
		},
		E_10: MsgInfoQueryAlertTypes{
			value: "10",
		},
		E_11: MsgInfoQueryAlertTypes{
			value: "11",
		},
		E_12: MsgInfoQueryAlertTypes{
			value: "12",
		},
		E_14: MsgInfoQueryAlertTypes{
			value: "14",
		},
		E_16: MsgInfoQueryAlertTypes{
			value: "16",
		},
		E_2: MsgInfoQueryAlertTypes{
			value: "2",
		},
		E_3: MsgInfoQueryAlertTypes{
			value: "3",
		},
		E_4: MsgInfoQueryAlertTypes{
			value: "4",
		},
		E_5: MsgInfoQueryAlertTypes{
			value: "5",
		},
		E_6: MsgInfoQueryAlertTypes{
			value: "6",
		},
		E_7: MsgInfoQueryAlertTypes{
			value: "7",
		},
		E_9: MsgInfoQueryAlertTypes{
			value: "9",
		},
	}
}

func (c MsgInfoQueryAlertTypes) Value() string {
	return c.value
}

func (c MsgInfoQueryAlertTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MsgInfoQueryAlertTypes) UnmarshalJSON(b []byte) error {
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

type MsgInfoQuerySendAlertType struct {
	value string
}

type MsgInfoQuerySendAlertTypeEnum struct {
	LEVEL_ALERT      MsgInfoQuerySendAlertType
	OPERATION_NOTICE MsgInfoQuerySendAlertType
	RECOVER_ALERT    MsgInfoQuerySendAlertType
	RESTRAINING      MsgInfoQuerySendAlertType
	RESTRAIN_ALERT   MsgInfoQuerySendAlertType
}

func GetMsgInfoQuerySendAlertTypeEnum() MsgInfoQuerySendAlertTypeEnum {
	return MsgInfoQuerySendAlertTypeEnum{
		LEVEL_ALERT: MsgInfoQuerySendAlertType{
			value: "LEVEL_ALERT",
		},
		OPERATION_NOTICE: MsgInfoQuerySendAlertType{
			value: "OPERATION_NOTICE",
		},
		RECOVER_ALERT: MsgInfoQuerySendAlertType{
			value: "RECOVER_ALERT",
		},
		RESTRAINING: MsgInfoQuerySendAlertType{
			value: "RESTRAINING",
		},
		RESTRAIN_ALERT: MsgInfoQuerySendAlertType{
			value: "RESTRAIN_ALERT",
		},
	}
}

func (c MsgInfoQuerySendAlertType) Value() string {
	return c.value
}

func (c MsgInfoQuerySendAlertType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MsgInfoQuerySendAlertType) UnmarshalJSON(b []byte) error {
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
