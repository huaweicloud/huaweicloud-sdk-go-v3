package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowTaskResponse Response Object
type ShowTaskResponse struct {

	// 任务编号
	TaskId *string `json:"task_id,omitempty"`

	// 任务运行状态，暂考虑取值仅为 Running/Failed/Successed
	Status *ShowTaskResponseStatus `json:"status,omitempty"`

	// 如果提交任务使用了input_enable参数，并且创建任务使用的是json格式非文件方式，该值为输入的字符串; 对应数据结构参见创建任务时的结构体
	InputJson *interface{} `json:"input_json,omitempty"`

	// 如果提交任务使用了input_enable参数，并且创建任务使用的是文件方式，该值为OBS对应的文件绝对路径
	InputUrl *string `json:"input_url,omitempty"`

	// 开始时间（UTC）
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间（UTC）
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 创建时间（UTC）
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 任务处理结果，json格式；每个子服务该对象结构不同，框架层不解析具体key，运行态直接拷贝算法服务返回信息、
	OutputJson *interface{} `json:"output_json,omitempty"`

	// 任务结果文件对应的绝对地址，具体值由租户OBS对应的待存储路径前缀和文件名组成，文件名服务端固定用task_id命名
	OutputUrl      *string `json:"output_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}

type ShowTaskResponseStatus struct {
	value string
}

type ShowTaskResponseStatusEnum struct {
	RUNNING   ShowTaskResponseStatus
	FAILED    ShowTaskResponseStatus
	SUCCESSED ShowTaskResponseStatus
}

func GetShowTaskResponseStatusEnum() ShowTaskResponseStatusEnum {
	return ShowTaskResponseStatusEnum{
		RUNNING: ShowTaskResponseStatus{
			value: "Running",
		},
		FAILED: ShowTaskResponseStatus{
			value: "Failed",
		},
		SUCCESSED: ShowTaskResponseStatus{
			value: "Successed",
		},
	}
}

func (c ShowTaskResponseStatus) Value() string {
	return c.value
}

func (c ShowTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
