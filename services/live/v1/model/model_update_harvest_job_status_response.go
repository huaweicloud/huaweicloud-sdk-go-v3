package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateHarvestJobStatusResponse Response Object
type UpdateHarvestJobStatusResponse struct {

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
	Status *UpdateHarvestJobStatusResponseStatus `json:"status,omitempty"`

	// 任务状态为FAILED时，最大允许重试的次数
	MaxRetryCnt *int32 `json:"max_retry_cnt,omitempty"`

	PackageInfo    *VodPackageInfo `json:"package_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateHarvestJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHarvestJobStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateHarvestJobStatusResponse", string(data)}, " ")
}

type UpdateHarvestJobStatusResponseStatus struct {
	value string
}

type UpdateHarvestJobStatusResponseStatusEnum struct {
	UNSTART     UpdateHarvestJobStatusResponseStatus
	IN_PROGRESS UpdateHarvestJobStatusResponseStatus
	SUCCEEDED   UpdateHarvestJobStatusResponseStatus
	FAILED      UpdateHarvestJobStatusResponseStatus
}

func GetUpdateHarvestJobStatusResponseStatusEnum() UpdateHarvestJobStatusResponseStatusEnum {
	return UpdateHarvestJobStatusResponseStatusEnum{
		UNSTART: UpdateHarvestJobStatusResponseStatus{
			value: "UNSTART",
		},
		IN_PROGRESS: UpdateHarvestJobStatusResponseStatus{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: UpdateHarvestJobStatusResponseStatus{
			value: "SUCCEEDED",
		},
		FAILED: UpdateHarvestJobStatusResponseStatus{
			value: "FAILED",
		},
	}
}

func (c UpdateHarvestJobStatusResponseStatus) Value() string {
	return c.value
}

func (c UpdateHarvestJobStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHarvestJobStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
