package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobInfo struct {
	Name *string `json:"name,omitempty"`

	Nodes *[]Node `json:"nodes,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	Params *[]JobParam `json:"params,omitempty"`

	Directory *string `json:"directory,omitempty"`

	JobType *JobInfoJobType `json:"jobType,omitempty"`

	BasicConfig *BasicInfo `json:"basicConfig,omitempty"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}

type JobInfoJobType struct {
	value string
}

type JobInfoJobTypeEnum struct {
	BATCH     JobInfoJobType
	REAL_TIME JobInfoJobType
}

func GetJobInfoJobTypeEnum() JobInfoJobTypeEnum {
	return JobInfoJobTypeEnum{
		BATCH: JobInfoJobType{
			value: "BATCH",
		},
		REAL_TIME: JobInfoJobType{
			value: "REAL_TIME",
		},
	}
}

func (c JobInfoJobType) Value() string {
	return c.value
}

func (c JobInfoJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoJobType) UnmarshalJSON(b []byte) error {
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
