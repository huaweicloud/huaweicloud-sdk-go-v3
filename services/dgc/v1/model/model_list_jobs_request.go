package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 分页参数：每页限定数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数：页数
	Offset *int32 `json:"offset,omitempty"`

	// 作业类型:  - REAL_TIME: 实时处理   - BATCH: 批处理
	JobType *ListJobsRequestJobType `json:"jobType,omitempty"`

	// 作业名称
	JobName *string `json:"jobName,omitempty"`

	// 作业标签
	Tags *string `json:"tags,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}

type ListJobsRequestJobType struct {
	value string
}

type ListJobsRequestJobTypeEnum struct {
	REAL_TIME ListJobsRequestJobType
	BATCH     ListJobsRequestJobType
}

func GetListJobsRequestJobTypeEnum() ListJobsRequestJobTypeEnum {
	return ListJobsRequestJobTypeEnum{
		REAL_TIME: ListJobsRequestJobType{
			value: "REAL_TIME",
		},
		BATCH: ListJobsRequestJobType{
			value: "BATCH",
		},
	}
}

func (c ListJobsRequestJobType) Value() string {
	return c.value
}

func (c ListJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestJobType) UnmarshalJSON(b []byte) error {
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
