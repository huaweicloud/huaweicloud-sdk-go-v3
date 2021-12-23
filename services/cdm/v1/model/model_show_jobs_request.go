package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowJobsRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`
	// 查询多个作业用all,查询单个作业输入作业名

	JobName string `json:"job_name"`
	// 当“job_name”为“all”时，此参数用于模糊过滤作业

	Filter *string `json:"filter,omitempty"`
	// 指定作业页号

	PageNo *int32 `json:"page_no,omitempty"`
	// 每页作业数，值在10-100之间

	PageSize *int32 `json:"page_size,omitempty"`
	// 作业类型: - jobType=NORMAL_JOB：表示查询表/文件迁移的作业。 - jobType=BATCH_JOB：表示查询整库迁移的作业。 - jobType=SCENARIO_JOB：表示查询场景迁移的作业。 - 不指定该参数时，默认只查询表/文件迁移的作业。

	JobType *ShowJobsRequestJobType `json:"jobType,omitempty"`
}

func (o ShowJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobsRequest struct{}"
	}

	return strings.Join([]string{"ShowJobsRequest", string(data)}, " ")
}

type ShowJobsRequestJobType struct {
	value string
}

type ShowJobsRequestJobTypeEnum struct {
	NORMAL_JOB   ShowJobsRequestJobType
	BATCH_JOB    ShowJobsRequestJobType
	SCENARIO_JOB ShowJobsRequestJobType
}

func GetShowJobsRequestJobTypeEnum() ShowJobsRequestJobTypeEnum {
	return ShowJobsRequestJobTypeEnum{
		NORMAL_JOB: ShowJobsRequestJobType{
			value: "NORMAL_JOB",
		},
		BATCH_JOB: ShowJobsRequestJobType{
			value: "BATCH_JOB",
		},
		SCENARIO_JOB: ShowJobsRequestJobType{
			value: "SCENARIO_JOB",
		},
	}
}

func (c ShowJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobsRequestJobType) UnmarshalJSON(b []byte) error {
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
