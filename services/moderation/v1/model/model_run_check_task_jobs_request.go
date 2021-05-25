package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type RunCheckTaskJobsRequest struct {
	// 图像内容审核任务处理状态如下：  - created 已创建  - running 正在处理  - finish 已完成  - failed 处理失败

	Status *RunCheckTaskJobsRequestStatus `json:"status,omitempty"`
}

func (o RunCheckTaskJobsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunCheckTaskJobsRequest struct{}"
	}

	return strings.Join([]string{"RunCheckTaskJobsRequest", string(data)}, " ")
}

type RunCheckTaskJobsRequestStatus struct {
	value string
}

type RunCheckTaskJobsRequestStatusEnum struct {
	CREATED RunCheckTaskJobsRequestStatus
	RUNNING RunCheckTaskJobsRequestStatus
	FINISH  RunCheckTaskJobsRequestStatus
	FAILED  RunCheckTaskJobsRequestStatus
}

func GetRunCheckTaskJobsRequestStatusEnum() RunCheckTaskJobsRequestStatusEnum {
	return RunCheckTaskJobsRequestStatusEnum{
		CREATED: RunCheckTaskJobsRequestStatus{
			value: "created",
		},
		RUNNING: RunCheckTaskJobsRequestStatus{
			value: "running",
		},
		FINISH: RunCheckTaskJobsRequestStatus{
			value: "finish",
		},
		FAILED: RunCheckTaskJobsRequestStatus{
			value: "failed",
		},
	}
}

func (c RunCheckTaskJobsRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *RunCheckTaskJobsRequestStatus) UnmarshalJSON(b []byte) error {
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
