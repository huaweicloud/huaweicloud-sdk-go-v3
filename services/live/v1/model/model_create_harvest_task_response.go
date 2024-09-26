package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHarvestTaskResponse Response Object
type CreateHarvestTaskResponse struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 频道推流域名
	Domain *string `json:"domain,omitempty"`

	// 组名或应用名
	AppName *string `json:"app_name,omitempty"`

	// 频道ID。频道唯一标识，为必填项。频道ID不建议输入下划线“_”，否则会影响转码和截图任务
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
	Status *CreateHarvestTaskResponseStatus `json:"status,omitempty"`

	// 任务状态为FAILED时，最大允许重试的次数
	MaxRetryCnt *int32 `json:"max_retry_cnt,omitempty"`

	PackageInfo    *VodPackageInfo `json:"package_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateHarvestTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHarvestTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateHarvestTaskResponse", string(data)}, " ")
}

type CreateHarvestTaskResponseStatus struct {
	value string
}

type CreateHarvestTaskResponseStatusEnum struct {
	UNSTART     CreateHarvestTaskResponseStatus
	IN_PROGRESS CreateHarvestTaskResponseStatus
	SUCCEEDED   CreateHarvestTaskResponseStatus
	FAILED      CreateHarvestTaskResponseStatus
}

func GetCreateHarvestTaskResponseStatusEnum() CreateHarvestTaskResponseStatusEnum {
	return CreateHarvestTaskResponseStatusEnum{
		UNSTART: CreateHarvestTaskResponseStatus{
			value: "UNSTART",
		},
		IN_PROGRESS: CreateHarvestTaskResponseStatus{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: CreateHarvestTaskResponseStatus{
			value: "SUCCEEDED",
		},
		FAILED: CreateHarvestTaskResponseStatus{
			value: "FAILED",
		},
	}
}

func (c CreateHarvestTaskResponseStatus) Value() string {
	return c.value
}

func (c CreateHarvestTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHarvestTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
