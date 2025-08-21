package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLatestPipelineJobsResponse Response Object
type ListLatestPipelineJobsResponse struct {

	// 流水线ID
	Id *int32 `json:"id,omitempty"`

	// 流水线状态，pending为排队，running为运行中，success为成功，failed为失败，canceled为取消，skipped为跳过，timedout为超时
	Status *ListLatestPipelineJobsResponseStatus `json:"status,omitempty"`

	Stages         *[]PipelineStageJobDto `json:"stages,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListLatestPipelineJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLatestPipelineJobsResponse struct{}"
	}

	return strings.Join([]string{"ListLatestPipelineJobsResponse", string(data)}, " ")
}

type ListLatestPipelineJobsResponseStatus struct {
	value string
}

type ListLatestPipelineJobsResponseStatusEnum struct {
	PENDING  ListLatestPipelineJobsResponseStatus
	RUNNING  ListLatestPipelineJobsResponseStatus
	SUCCESS  ListLatestPipelineJobsResponseStatus
	FAILED   ListLatestPipelineJobsResponseStatus
	CANCELED ListLatestPipelineJobsResponseStatus
	SKIPPED  ListLatestPipelineJobsResponseStatus
	TIMEDOUT ListLatestPipelineJobsResponseStatus
}

func GetListLatestPipelineJobsResponseStatusEnum() ListLatestPipelineJobsResponseStatusEnum {
	return ListLatestPipelineJobsResponseStatusEnum{
		PENDING: ListLatestPipelineJobsResponseStatus{
			value: "pending",
		},
		RUNNING: ListLatestPipelineJobsResponseStatus{
			value: "running",
		},
		SUCCESS: ListLatestPipelineJobsResponseStatus{
			value: "success",
		},
		FAILED: ListLatestPipelineJobsResponseStatus{
			value: "failed",
		},
		CANCELED: ListLatestPipelineJobsResponseStatus{
			value: "canceled",
		},
		SKIPPED: ListLatestPipelineJobsResponseStatus{
			value: "skipped",
		},
		TIMEDOUT: ListLatestPipelineJobsResponseStatus{
			value: "timedout",
		},
	}
}

func (c ListLatestPipelineJobsResponseStatus) Value() string {
	return c.value
}

func (c ListLatestPipelineJobsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLatestPipelineJobsResponseStatus) UnmarshalJSON(b []byte) error {
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
