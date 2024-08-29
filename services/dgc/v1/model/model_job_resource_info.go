package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobResourceInfo struct {

	// 资源名
	Name string `json:"name"`

	// 资源类型
	Type JobResourceInfoType `json:"type"`

	// 替换后的资源名
	Replace *string `json:"replace,omitempty"`
}

func (o JobResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResourceInfo struct{}"
	}

	return strings.Join([]string{"JobResourceInfo", string(data)}, " ")
}

type JobResourceInfoType struct {
	value string
}

type JobResourceInfoTypeEnum struct {
	DWS_CONNECTION       JobResourceInfoType
	DIS_STREAM           JobResourceInfoType
	CDM_CLUSTER          JobResourceInfoType
	MRS_CLUSTER          JobResourceInfoType
	MRS_HIVE_CONNECTION  JobResourceInfoType
	MRS_SPARK_CONNECTION JobResourceInfoType
	GES_GRAPH            JobResourceInfoType
	ROMA_INSTANCE        JobResourceInfoType
	TICS_INSTANCE        JobResourceInfoType
	DRS_TASK             JobResourceInfoType
}

func GetJobResourceInfoTypeEnum() JobResourceInfoTypeEnum {
	return JobResourceInfoTypeEnum{
		DWS_CONNECTION: JobResourceInfoType{
			value: "DWS_CONNECTION",
		},
		DIS_STREAM: JobResourceInfoType{
			value: "DIS_STREAM",
		},
		CDM_CLUSTER: JobResourceInfoType{
			value: "CDM_CLUSTER",
		},
		MRS_CLUSTER: JobResourceInfoType{
			value: "MRS_CLUSTER",
		},
		MRS_HIVE_CONNECTION: JobResourceInfoType{
			value: "MRS_HIVE_CONNECTION",
		},
		MRS_SPARK_CONNECTION: JobResourceInfoType{
			value: "MRS_SPARK_CONNECTION",
		},
		GES_GRAPH: JobResourceInfoType{
			value: "GES_GRAPH",
		},
		ROMA_INSTANCE: JobResourceInfoType{
			value: "ROMA_INSTANCE",
		},
		TICS_INSTANCE: JobResourceInfoType{
			value: "TICS_INSTANCE",
		},
		DRS_TASK: JobResourceInfoType{
			value: "DRS_TASK",
		},
	}
}

func (c JobResourceInfoType) Value() string {
	return c.value
}

func (c JobResourceInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobResourceInfoType) UnmarshalJSON(b []byte) error {
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
