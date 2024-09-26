package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HarvestTaskCreateSucRsp 创建成功响应
type HarvestTaskCreateSucRsp struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 频道推流域名
	Domain string `json:"domain"`

	// 组名或应用名
	AppName string `json:"app_name"`

	// 频道ID。频道唯一标识，为必填项。频道ID不建议输入下划线“_”，否则会影响转码和截图任务
	Id string `json:"id"`

	// 开始时间。Unix时间错，单位为秒
	StartTime int32 `json:"start_time"`

	// 结束时间。Unix时间错，单位为秒
	EndTime int32 `json:"end_time"`

	// 任务创建时间。Unix时间错，单位为秒
	CreateTime *int32 `json:"create_time,omitempty"`

	// 事件名称
	EventName string `json:"event_name"`

	// 任务描述
	TaskDesc string `json:"task_desc"`

	// 任务状态，取值为 [ UNSTART、IN_PROGRESS、SUCCEEDED、FAILED ]
	Status HarvestTaskCreateSucRspStatus `json:"status"`

	// 任务状态为FAILED时，最大允许重试的次数
	MaxRetryCnt int32 `json:"max_retry_cnt"`

	PackageInfo *VodPackageInfo `json:"package_info,omitempty"`
}

func (o HarvestTaskCreateSucRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HarvestTaskCreateSucRsp struct{}"
	}

	return strings.Join([]string{"HarvestTaskCreateSucRsp", string(data)}, " ")
}

type HarvestTaskCreateSucRspStatus struct {
	value string
}

type HarvestTaskCreateSucRspStatusEnum struct {
	UNSTART     HarvestTaskCreateSucRspStatus
	IN_PROGRESS HarvestTaskCreateSucRspStatus
	SUCCEEDED   HarvestTaskCreateSucRspStatus
	FAILED      HarvestTaskCreateSucRspStatus
}

func GetHarvestTaskCreateSucRspStatusEnum() HarvestTaskCreateSucRspStatusEnum {
	return HarvestTaskCreateSucRspStatusEnum{
		UNSTART: HarvestTaskCreateSucRspStatus{
			value: "UNSTART",
		},
		IN_PROGRESS: HarvestTaskCreateSucRspStatus{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: HarvestTaskCreateSucRspStatus{
			value: "SUCCEEDED",
		},
		FAILED: HarvestTaskCreateSucRspStatus{
			value: "FAILED",
		},
	}
}

func (c HarvestTaskCreateSucRspStatus) Value() string {
	return c.value
}

func (c HarvestTaskCreateSucRspStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HarvestTaskCreateSucRspStatus) UnmarshalJSON(b []byte) error {
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
