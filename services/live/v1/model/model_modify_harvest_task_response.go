package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyHarvestTaskResponse Response Object
type ModifyHarvestTaskResponse struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 频道推流域名
	Domain *string `json:"domain,omitempty"`

	// 组名或应用名
	AppName *string `json:"app_name,omitempty"`

	// 频道ID。频道唯一标识，为必填项。
	Id *string `json:"id,omitempty"`

	// 开始时间。Unix时间错，单位为秒
	StartTime *int32 `json:"start_time,omitempty"`

	// 结束时间。Unix时间错，单位为秒
	EndTime *int32 `json:"end_time,omitempty"`

	// 任务创建时间。Unix时间错，单位为秒
	CreateTime *int32 `json:"create_time,omitempty"`

	// 事件名称
	EventName *string `json:"event_name,omitempty"`

	// 任务描述
	TaskDesc *string `json:"task_desc,omitempty"`

	// 任务状态，取值为 [ UNSTART、IN_PROGRESS、SUCCEEDED、FAILED ]
	Status *ModifyHarvestTaskResponseStatus `json:"status,omitempty"`

	// 任务状态为FAILED时，最大允许重试的次数
	MaxRetryCnt *int32 `json:"max_retry_cnt,omitempty"`

	PackageInfo    *VodPackageInfo `json:"package_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ModifyHarvestTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHarvestTaskResponse struct{}"
	}

	return strings.Join([]string{"ModifyHarvestTaskResponse", string(data)}, " ")
}

type ModifyHarvestTaskResponseStatus struct {
	value string
}

type ModifyHarvestTaskResponseStatusEnum struct {
	UNSTART     ModifyHarvestTaskResponseStatus
	IN_PROGRESS ModifyHarvestTaskResponseStatus
	SUCCEEDED   ModifyHarvestTaskResponseStatus
	FAILED      ModifyHarvestTaskResponseStatus
}

func GetModifyHarvestTaskResponseStatusEnum() ModifyHarvestTaskResponseStatusEnum {
	return ModifyHarvestTaskResponseStatusEnum{
		UNSTART: ModifyHarvestTaskResponseStatus{
			value: "UNSTART",
		},
		IN_PROGRESS: ModifyHarvestTaskResponseStatus{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: ModifyHarvestTaskResponseStatus{
			value: "SUCCEEDED",
		},
		FAILED: ModifyHarvestTaskResponseStatus{
			value: "FAILED",
		},
	}
}

func (c ModifyHarvestTaskResponseStatus) Value() string {
	return c.value
}

func (c ModifyHarvestTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyHarvestTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
