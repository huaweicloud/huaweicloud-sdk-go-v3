package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RunCheckTaskJobsRequest Request Object
type RunCheckTaskJobsRequest struct {

	// 图像内容审核任务处理状态如下：  - created 已创建  - running 正在处理  - finish 已完成  - failed 处理失败
	Status *RunCheckTaskJobsRequestStatus `json:"status,omitempty"`

	// 偏移量， 默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，默认为符合查询条件的总任务数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o RunCheckTaskJobsRequest) String() string {
	data, err := utils.Marshal(o)
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

func (c RunCheckTaskJobsRequestStatus) Value() string {
	return c.value
}

func (c RunCheckTaskJobsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunCheckTaskJobsRequestStatus) UnmarshalJSON(b []byte) error {
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
