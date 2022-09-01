package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ThumbTask struct {

	// 任务ID。  截图服务接受任务后产生的任务ID。一次最多10个。
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务状态
	Status *ThumbTaskStatus `json:"status,omitempty" xml:"status"`

	// 起始时间。  格式为yyyymmddhhmmss。必须是与时区无关的UTC时间，指定task_id时该参数无效。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 结束时间。  格式为yyyymmddhhmmss。必须是与时区无关的UTC时间，指定task_id时该参数无效。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	// 输出文件名。
	OutputFileName *string `json:"output_file_name,omitempty" xml:"output_file_name"`

	// 用户数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	// 任务描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 截图文件信息。
	ThumbnailInfo *[]PicInfo `json:"thumbnail_info,omitempty" xml:"thumbnail_info"`
}

func (o ThumbTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThumbTask struct{}"
	}

	return strings.Join([]string{"ThumbTask", string(data)}, " ")
}

type ThumbTaskStatus struct {
	value string
}

type ThumbTaskStatusEnum struct {
	NO_TASK    ThumbTaskStatus
	WAITING    ThumbTaskStatus
	PROCESSING ThumbTaskStatus
	SUCCEEDED  ThumbTaskStatus
	FAILED     ThumbTaskStatus
	CANCELED   ThumbTaskStatus
}

func GetThumbTaskStatusEnum() ThumbTaskStatusEnum {
	return ThumbTaskStatusEnum{
		NO_TASK: ThumbTaskStatus{
			value: "NO_TASK",
		},
		WAITING: ThumbTaskStatus{
			value: "WAITING",
		},
		PROCESSING: ThumbTaskStatus{
			value: "PROCESSING",
		},
		SUCCEEDED: ThumbTaskStatus{
			value: "SUCCEEDED",
		},
		FAILED: ThumbTaskStatus{
			value: "FAILED",
		},
		CANCELED: ThumbTaskStatus{
			value: "CANCELED",
		},
	}
}

func (c ThumbTaskStatus) Value() string {
	return c.value
}

func (c ThumbTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThumbTaskStatus) UnmarshalJSON(b []byte) error {
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
