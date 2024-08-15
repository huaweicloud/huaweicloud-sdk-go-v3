package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobScriptOrderInfoModel 基本信息
type JobScriptOrderInfoModel struct {

	// 执行uuid
	ExecuteUuid *string `json:"execute_uuid,omitempty"`

	// 执行创建时间
	GmtCreated *int64 `json:"gmt_created,omitempty"`

	// 执行完成时间
	GmtFinished *int64 `json:"gmt_finished,omitempty"`

	// 执行耗时，单位：秒
	ExecuteCosts *int64 `json:"execute_costs,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 执行状态
	Status *JobScriptOrderInfoModelStatus `json:"status,omitempty"`

	Properties *JobScriptOrderInfoProp `json:"properties,omitempty"`
}

func (o JobScriptOrderInfoModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptOrderInfoModel struct{}"
	}

	return strings.Join([]string{"JobScriptOrderInfoModel", string(data)}, " ")
}

type JobScriptOrderInfoModelStatus struct {
	value string
}

type JobScriptOrderInfoModelStatusEnum struct {
	READY      JobScriptOrderInfoModelStatus
	PROCESSING JobScriptOrderInfoModelStatus
	ABNORMAL   JobScriptOrderInfoModelStatus
	PAUSED     JobScriptOrderInfoModelStatus
	CANCELED   JobScriptOrderInfoModelStatus
	FINISHED   JobScriptOrderInfoModelStatus
}

func GetJobScriptOrderInfoModelStatusEnum() JobScriptOrderInfoModelStatusEnum {
	return JobScriptOrderInfoModelStatusEnum{
		READY: JobScriptOrderInfoModelStatus{
			value: "READY",
		},
		PROCESSING: JobScriptOrderInfoModelStatus{
			value: "PROCESSING",
		},
		ABNORMAL: JobScriptOrderInfoModelStatus{
			value: "ABNORMAL",
		},
		PAUSED: JobScriptOrderInfoModelStatus{
			value: "PAUSED",
		},
		CANCELED: JobScriptOrderInfoModelStatus{
			value: "CANCELED",
		},
		FINISHED: JobScriptOrderInfoModelStatus{
			value: "FINISHED",
		},
	}
}

func (c JobScriptOrderInfoModelStatus) Value() string {
	return c.value
}

func (c JobScriptOrderInfoModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobScriptOrderInfoModelStatus) UnmarshalJSON(b []byte) error {
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
