package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobScriptOrderListModel 作业脚本工单列表
type JobScriptOrderListModel struct {

	// 主键id，对应job_order_do的主键
	OrderId int64 `json:"order_id"`

	// 工单名称
	OrderName string `json:"order_name"`

	// 列表跳转到详情时，用这个uuid，对应execute_data_do的execute_uuid
	ExecuteUuid string `json:"execute_uuid"`

	// 创建时间
	GmtCreated *int64 `json:"gmt_created,omitempty"`

	// 完成时间
	GmtFinished *int64 `json:"gmt_finished,omitempty"`

	// 执行耗时，单位：秒
	ExecuteCosts *int64 `json:"execute_costs,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 工单状态
	Status *JobScriptOrderListModelStatus `json:"status,omitempty"`

	Properties *JobScriptOrderListProp `json:"properties,omitempty"`
}

func (o JobScriptOrderListModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptOrderListModel struct{}"
	}

	return strings.Join([]string{"JobScriptOrderListModel", string(data)}, " ")
}

type JobScriptOrderListModelStatus struct {
	value string
}

type JobScriptOrderListModelStatusEnum struct {
	READY      JobScriptOrderListModelStatus
	PROCESSING JobScriptOrderListModelStatus
	ABNORMAL   JobScriptOrderListModelStatus
	PAUSED     JobScriptOrderListModelStatus
	CANCELED   JobScriptOrderListModelStatus
	FINISHED   JobScriptOrderListModelStatus
}

func GetJobScriptOrderListModelStatusEnum() JobScriptOrderListModelStatusEnum {
	return JobScriptOrderListModelStatusEnum{
		READY: JobScriptOrderListModelStatus{
			value: "READY",
		},
		PROCESSING: JobScriptOrderListModelStatus{
			value: "PROCESSING",
		},
		ABNORMAL: JobScriptOrderListModelStatus{
			value: "ABNORMAL",
		},
		PAUSED: JobScriptOrderListModelStatus{
			value: "PAUSED",
		},
		CANCELED: JobScriptOrderListModelStatus{
			value: "CANCELED",
		},
		FINISHED: JobScriptOrderListModelStatus{
			value: "FINISHED",
		},
	}
}

func (c JobScriptOrderListModelStatus) Value() string {
	return c.value
}

func (c JobScriptOrderListModelStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobScriptOrderListModelStatus) UnmarshalJSON(b []byte) error {
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
