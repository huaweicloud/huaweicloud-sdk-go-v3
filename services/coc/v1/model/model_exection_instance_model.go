package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExectionInstanceModel 执行主机实例详情
type ExectionInstanceModel struct {

	// 主键id
	Id *int64 `json:"id,omitempty"`

	TargetInstance *ResourceInstance `json:"target_instance,omitempty"`

	// 创建时间
	GmtCreated *int64 `json:"gmt_created,omitempty"`

	// 完成时间
	GmtFinished *int64 `json:"gmt_finished,omitempty"`

	// 耗时时间，单位:秒
	ExecuteCosts *int64 `json:"execute_costs,omitempty"`

	// 实例执行状态
	Status *ExectionInstanceModelStatus `json:"status,omitempty"`

	// 实例执行日志
	Message *string `json:"message,omitempty"`
}

func (o ExectionInstanceModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExectionInstanceModel struct{}"
	}

	return strings.Join([]string{"ExectionInstanceModel", string(data)}, " ")
}

type ExectionInstanceModelStatus struct {
	value string
}

type ExectionInstanceModelStatusEnum struct {
	READY      ExectionInstanceModelStatus
	PROCESSING ExectionInstanceModelStatus
	ABNORMAL   ExectionInstanceModelStatus
	CANCELED   ExectionInstanceModelStatus
	FINISHED   ExectionInstanceModelStatus
	ROLLBACKED ExectionInstanceModelStatus
}

func GetExectionInstanceModelStatusEnum() ExectionInstanceModelStatusEnum {
	return ExectionInstanceModelStatusEnum{
		READY: ExectionInstanceModelStatus{
			value: "READY",
		},
		PROCESSING: ExectionInstanceModelStatus{
			value: "PROCESSING",
		},
		ABNORMAL: ExectionInstanceModelStatus{
			value: "ABNORMAL",
		},
		CANCELED: ExectionInstanceModelStatus{
			value: "CANCELED",
		},
		FINISHED: ExectionInstanceModelStatus{
			value: "FINISHED",
		},
		ROLLBACKED: ExectionInstanceModelStatus{
			value: "ROLLBACKED",
		},
	}
}

func (c ExectionInstanceModelStatus) Value() string {
	return c.value
}

func (c ExectionInstanceModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExectionInstanceModelStatus) UnmarshalJSON(b []byte) error {
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
