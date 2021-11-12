package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 作业的输入配置，必填。
type TaskInput struct {
	// 作业的输入类型，必填。可选类型有obs（对象存储服务存储的文件），vis（视频接入服务的视频流），url（指定的文件地址或取流地址），edgecamera（绑定在IEF的边缘摄像头），edgerestful（从自定义的流媒体服务器通过Restful接口获取取流地址），vcn（VCN设备）。

	Type TaskInputType `json:"type"`
	// 作业的输入详情，必填。针对不同的输入类型有不同的配置。创建时允许填写多路输入，但更新时只允许填写一路输入。

	Data []TaskInputData `json:"data"`

	Vcn *TaskInputVcn `json:"vcn,omitempty"`
}

func (o TaskInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInput struct{}"
	}

	return strings.Join([]string{"TaskInput", string(data)}, " ")
}

type TaskInputType struct {
	value string
}

type TaskInputTypeEnum struct {
	OBS         TaskInputType
	VIS         TaskInputType
	URL         TaskInputType
	EDGECAMERA  TaskInputType
	EDGERESTFUL TaskInputType
	VCN         TaskInputType
}

func GetTaskInputTypeEnum() TaskInputTypeEnum {
	return TaskInputTypeEnum{
		OBS: TaskInputType{
			value: "obs",
		},
		VIS: TaskInputType{
			value: "vis",
		},
		URL: TaskInputType{
			value: "url",
		},
		EDGECAMERA: TaskInputType{
			value: "edgecamera",
		},
		EDGERESTFUL: TaskInputType{
			value: "edgerestful",
		},
		VCN: TaskInputType{
			value: "vcn",
		},
	}
}

func (c TaskInputType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskInputType) UnmarshalJSON(b []byte) error {
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
