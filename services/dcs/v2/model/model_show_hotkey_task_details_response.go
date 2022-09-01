package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowHotkeyTaskDetailsResponse struct {

	// 热key分析记录ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 分析任务状态
	Status *ShowHotkeyTaskDetailsResponseStatus `json:"status,omitempty" xml:"status"`

	// 分析方式
	ScanType *ShowHotkeyTaskDetailsResponseScanType `json:"scan_type,omitempty" xml:"scan_type"`

	// 分析任务创建时间,格式为：\"2020-06-15T02:21:18.669Z\"
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 分析任务开始时间,格式为：\"2020-06-15T02:21:18.669Z\"
	StartedAt *string `json:"started_at,omitempty" xml:"started_at"`

	// 分析任务结束时间,格式为：\"2020-06-15T02:21:18.669Z\"
	FinishedAt *string `json:"finished_at,omitempty" xml:"finished_at"`

	// 热key的数量
	Num *int32 `json:"num,omitempty" xml:"num"`

	// 热key记录
	Keys           *[]HotkeysBody `json:"keys,omitempty" xml:"keys"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowHotkeyTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotkeyTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowHotkeyTaskDetailsResponse", string(data)}, " ")
}

type ShowHotkeyTaskDetailsResponseStatus struct {
	value string
}

type ShowHotkeyTaskDetailsResponseStatusEnum struct {
	WAITING ShowHotkeyTaskDetailsResponseStatus
	RUNNING ShowHotkeyTaskDetailsResponseStatus
	SUCCESS ShowHotkeyTaskDetailsResponseStatus
	FAILED  ShowHotkeyTaskDetailsResponseStatus
}

func GetShowHotkeyTaskDetailsResponseStatusEnum() ShowHotkeyTaskDetailsResponseStatusEnum {
	return ShowHotkeyTaskDetailsResponseStatusEnum{
		WAITING: ShowHotkeyTaskDetailsResponseStatus{
			value: "waiting",
		},
		RUNNING: ShowHotkeyTaskDetailsResponseStatus{
			value: "running",
		},
		SUCCESS: ShowHotkeyTaskDetailsResponseStatus{
			value: "success",
		},
		FAILED: ShowHotkeyTaskDetailsResponseStatus{
			value: "failed",
		},
	}
}

func (c ShowHotkeyTaskDetailsResponseStatus) Value() string {
	return c.value
}

func (c ShowHotkeyTaskDetailsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHotkeyTaskDetailsResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowHotkeyTaskDetailsResponseScanType struct {
	value string
}

type ShowHotkeyTaskDetailsResponseScanTypeEnum struct {
	MANUAL ShowHotkeyTaskDetailsResponseScanType
	AUTO   ShowHotkeyTaskDetailsResponseScanType
}

func GetShowHotkeyTaskDetailsResponseScanTypeEnum() ShowHotkeyTaskDetailsResponseScanTypeEnum {
	return ShowHotkeyTaskDetailsResponseScanTypeEnum{
		MANUAL: ShowHotkeyTaskDetailsResponseScanType{
			value: "manual",
		},
		AUTO: ShowHotkeyTaskDetailsResponseScanType{
			value: "auto",
		},
	}
}

func (c ShowHotkeyTaskDetailsResponseScanType) Value() string {
	return c.value
}

func (c ShowHotkeyTaskDetailsResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHotkeyTaskDetailsResponseScanType) UnmarshalJSON(b []byte) error {
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
