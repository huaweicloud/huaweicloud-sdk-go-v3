package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFactoryJobsRequest Request Object
type ListFactoryJobsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 分页参数：每页限定数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数：页数
	Offset *int32 `json:"offset,omitempty"`

	// 作业类型:  - REAL_TIME: 实时处理  - BATCH: 批处理
	JobType *ListFactoryJobsRequestJobType `json:"job_type,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 作业标签
	Tags *string `json:"tags,omitempty"`
}

func (o ListFactoryJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryJobsRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryJobsRequest", string(data)}, " ")
}

type ListFactoryJobsRequestJobType struct {
	value string
}

type ListFactoryJobsRequestJobTypeEnum struct {
	REAL_TIME ListFactoryJobsRequestJobType
	BATCH     ListFactoryJobsRequestJobType
}

func GetListFactoryJobsRequestJobTypeEnum() ListFactoryJobsRequestJobTypeEnum {
	return ListFactoryJobsRequestJobTypeEnum{
		REAL_TIME: ListFactoryJobsRequestJobType{
			value: "REAL_TIME",
		},
		BATCH: ListFactoryJobsRequestJobType{
			value: "BATCH",
		},
	}
}

func (c ListFactoryJobsRequestJobType) Value() string {
	return c.value
}

func (c ListFactoryJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactoryJobsRequestJobType) UnmarshalJSON(b []byte) error {
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
