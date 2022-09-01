package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowBigkeyScanTaskDetailsResponse struct {

	// 大key分析记录ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 分析任务状态
	Status *ShowBigkeyScanTaskDetailsResponseStatus `json:"status,omitempty" xml:"status"`

	// 分析方式
	ScanType *ShowBigkeyScanTaskDetailsResponseScanType `json:"scan_type,omitempty" xml:"scan_type"`

	// 分析任务创建时间,格式为：\"2020-06-15T02:21:18.669Z\"
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 分析任务开始时间,格式为：\"2020-06-15T02:21:18.669Z\"
	StartedAt *string `json:"started_at,omitempty" xml:"started_at"`

	// 分析任务结束时间,格式为：\"2020-06-15T02:21:18.669Z\"
	FinishedAt *string `json:"finished_at,omitempty" xml:"finished_at"`

	// 大key的数量
	Num *int32 `json:"num,omitempty" xml:"num"`

	// 大key记录
	Keys           *[]BigkeysBody `json:"keys,omitempty" xml:"keys"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowBigkeyScanTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBigkeyScanTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowBigkeyScanTaskDetailsResponse", string(data)}, " ")
}

type ShowBigkeyScanTaskDetailsResponseStatus struct {
	value string
}

type ShowBigkeyScanTaskDetailsResponseStatusEnum struct {
	WAITING ShowBigkeyScanTaskDetailsResponseStatus
	RUNNING ShowBigkeyScanTaskDetailsResponseStatus
	SUCCESS ShowBigkeyScanTaskDetailsResponseStatus
	FAILED  ShowBigkeyScanTaskDetailsResponseStatus
}

func GetShowBigkeyScanTaskDetailsResponseStatusEnum() ShowBigkeyScanTaskDetailsResponseStatusEnum {
	return ShowBigkeyScanTaskDetailsResponseStatusEnum{
		WAITING: ShowBigkeyScanTaskDetailsResponseStatus{
			value: "waiting",
		},
		RUNNING: ShowBigkeyScanTaskDetailsResponseStatus{
			value: "running",
		},
		SUCCESS: ShowBigkeyScanTaskDetailsResponseStatus{
			value: "success",
		},
		FAILED: ShowBigkeyScanTaskDetailsResponseStatus{
			value: "failed",
		},
	}
}

func (c ShowBigkeyScanTaskDetailsResponseStatus) Value() string {
	return c.value
}

func (c ShowBigkeyScanTaskDetailsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBigkeyScanTaskDetailsResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowBigkeyScanTaskDetailsResponseScanType struct {
	value string
}

type ShowBigkeyScanTaskDetailsResponseScanTypeEnum struct {
	MANUAL ShowBigkeyScanTaskDetailsResponseScanType
	AUTO   ShowBigkeyScanTaskDetailsResponseScanType
}

func GetShowBigkeyScanTaskDetailsResponseScanTypeEnum() ShowBigkeyScanTaskDetailsResponseScanTypeEnum {
	return ShowBigkeyScanTaskDetailsResponseScanTypeEnum{
		MANUAL: ShowBigkeyScanTaskDetailsResponseScanType{
			value: "manual",
		},
		AUTO: ShowBigkeyScanTaskDetailsResponseScanType{
			value: "auto",
		},
	}
}

func (c ShowBigkeyScanTaskDetailsResponseScanType) Value() string {
	return c.value
}

func (c ShowBigkeyScanTaskDetailsResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBigkeyScanTaskDetailsResponseScanType) UnmarshalJSON(b []byte) error {
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
