package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobInfo struct {
	Name *string `json:"name,omitempty" xml:"name"`

	Nodes *[]Node `json:"nodes,omitempty" xml:"nodes"`

	Schedule *Schedule `json:"schedule,omitempty" xml:"schedule"`

	Params *[]JobParam `json:"params,omitempty" xml:"params"`

	Directory *string `json:"directory,omitempty" xml:"directory"`

	JobType *JobInfoJobType `json:"jobType,omitempty" xml:"jobType"`

	BasicConfig *BasicInfo `json:"basicConfig,omitempty" xml:"basicConfig"`
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
