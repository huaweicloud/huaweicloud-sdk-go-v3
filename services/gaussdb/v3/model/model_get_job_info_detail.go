package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 任务信息。
type GetJobInfoDetail struct {

	// 任务ID。
	Id string `json:"id" xml:"id"`

	// 任务名称。
	Name string `json:"name" xml:"name"`

	// 任务执行状态。  取值： - 值为“Running”，表示任务正在执行。 - 值为“Completed”，表示任务执行成功。 - 值为“Failed”，表示任务执行失败。
	Status GetJobInfoDetailStatus `json:"status" xml:"status"`

	// 创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为   +0800 说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	Created string `json:"created" xml:"created"`

	// 结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为   +0800 说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	Ended *string `json:"ended,omitempty" xml:"ended"`

	// 任务执行进度。执行中状态才返回执行进度，例如60%，否则返回\"\"。
	Process *string `json:"process,omitempty" xml:"process"`

	Instance *GetJobInstanceInfoDetail `json:"instance" xml:"instance"`

	Entities *GetJobEntitiesInfoDetail `json:"entities,omitempty" xml:"entities"`

	// 任务执行失败时的错误信息。
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason"`
}

func (o GetJobInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetJobInfoDetail struct{}"
	}

	return strings.Join([]string{"GetJobInfoDetail", string(data)}, " ")
}

type GetJobInfoDetailStatus struct {
	value string
}

type GetJobInfoDetailStatusEnum struct {
	RUNNING   GetJobInfoDetailStatus
	COMPLETED GetJobInfoDetailStatus
	FAILED    GetJobInfoDetailStatus
}

func GetGetJobInfoDetailStatusEnum() GetJobInfoDetailStatusEnum {
	return GetJobInfoDetailStatusEnum{
		RUNNING: GetJobInfoDetailStatus{
			value: "Running",
		},
		COMPLETED: GetJobInfoDetailStatus{
			value: "Completed",
		},
		FAILED: GetJobInfoDetailStatus{
			value: "Failed",
		},
	}
}

func (c GetJobInfoDetailStatus) Value() string {
	return c.value
}

func (c GetJobInfoDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetJobInfoDetailStatus) UnmarshalJSON(b []byte) error {
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
