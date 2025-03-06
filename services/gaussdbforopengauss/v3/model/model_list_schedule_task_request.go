package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScheduleTaskRequest Request Object
type ListScheduleTaskRequest struct {

	// 语言。
	XLanguage *ListScheduleTaskRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务状态。
	Status *ListScheduleTaskRequestStatus `json:"status,omitempty"`

	// 任务名称。
	Name *ListScheduleTaskRequestName `json:"name,omitempty"`

	// 开始时间，格式为yyyy-mm-ddThh:mm:ssZ。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式为yyyy-mm-ddThh:mm:ssZ。
	EndTime *string `json:"end_time,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"ListScheduleTaskRequest", string(data)}, " ")
}

type ListScheduleTaskRequestXLanguage struct {
	value string
}

type ListScheduleTaskRequestXLanguageEnum struct {
	ZH_CN ListScheduleTaskRequestXLanguage
	EN_US ListScheduleTaskRequestXLanguage
}

func GetListScheduleTaskRequestXLanguageEnum() ListScheduleTaskRequestXLanguageEnum {
	return ListScheduleTaskRequestXLanguageEnum{
		ZH_CN: ListScheduleTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListScheduleTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListScheduleTaskRequestXLanguage) Value() string {
	return c.value
}

func (c ListScheduleTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduleTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListScheduleTaskRequestStatus struct {
	value string
}

type ListScheduleTaskRequestStatusEnum struct {
	RUNNING   ListScheduleTaskRequestStatus
	COMPLETED ListScheduleTaskRequestStatus
	FAILED    ListScheduleTaskRequestStatus
	CANCELED  ListScheduleTaskRequestStatus
	PENDING   ListScheduleTaskRequestStatus
}

func GetListScheduleTaskRequestStatusEnum() ListScheduleTaskRequestStatusEnum {
	return ListScheduleTaskRequestStatusEnum{
		RUNNING: ListScheduleTaskRequestStatus{
			value: "Running",
		},
		COMPLETED: ListScheduleTaskRequestStatus{
			value: "Completed",
		},
		FAILED: ListScheduleTaskRequestStatus{
			value: "Failed",
		},
		CANCELED: ListScheduleTaskRequestStatus{
			value: "Canceled",
		},
		PENDING: ListScheduleTaskRequestStatus{
			value: "Pending",
		},
	}
}

func (c ListScheduleTaskRequestStatus) Value() string {
	return c.value
}

func (c ListScheduleTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduleTaskRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListScheduleTaskRequestName struct {
	value string
}

type ListScheduleTaskRequestNameEnum struct {
	HOTFIX_VERSION_UPGRADE ListScheduleTaskRequestName
}

func GetListScheduleTaskRequestNameEnum() ListScheduleTaskRequestNameEnum {
	return ListScheduleTaskRequestNameEnum{
		HOTFIX_VERSION_UPGRADE: ListScheduleTaskRequestName{
			value: "HOTFIX_VERSION_UPGRADE",
		},
	}
}

func (c ListScheduleTaskRequestName) Value() string {
	return c.value
}

func (c ListScheduleTaskRequestName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduleTaskRequestName) UnmarshalJSON(b []byte) error {
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
