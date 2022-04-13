package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type UpdateTrackerRequestBody struct {
	// 标识追踪器类型。 目前支持系统追踪器类型有管理类追踪器(system)和数据类追踪器(data)。 数据类追踪器和管理类追踪器共同参数有：is_lts_enabled, obs_info; 管理类追踪器参数：is_support_trace_files_encryption, kms_id, is_support_validate, is_support_validate; 数据类追踪器参数：tracker_name, data_bucket。

	TrackerType UpdateTrackerRequestBodyTrackerType `json:"tracker_type"`
	// 标识追踪器名称。 当\"tracker_type\"参数值为\"system\"时该参数为默认值\"system\"。 当\"tracker_type\"参数值为\"data\"时该参数需要指定追踪器名称\"。

	TrackerName string `json:"tracker_name"`
	// 标识追踪器状态，该接口中可修改的状态包括正常（enabled）和停止（disabled）。如果选择修改状态为停止，则修改成功后追踪器停止记录事件。

	Status *UpdateTrackerRequestBodyStatus `json:"status,omitempty"`
	// 是否打开事件分析。

	IsLtsEnabled *bool `json:"is_lts_enabled,omitempty"`

	ObsInfo *TrackerObsInfo `json:"obs_info,omitempty"`
	// 事件文件转储加密功能开关。 当\"tracker_type\"参数值为\"system\"时该参数值有效。 该参数必须与kms_id参数同时使用。

	IsSupportTraceFilesEncryption *bool `json:"is_support_trace_files_encryption,omitempty"`
	// 事件文件转储加密所采用的秘钥id（从KMS获取）。 当\"tracker_type\"参数值为\"system\"时该参数值有效。 当\"is_support_trace_files_encryption\"参数值为“是”时，此参数为必选项。

	KmsId *string `json:"kms_id,omitempty"`
	// 事件文件转储时是否打开事件文件校验。 当\"tracker_type\"参数值为\"system\"时该参数值有效。

	IsSupportValidate *bool `json:"is_support_validate,omitempty"`

	DataBucket *DataBucket `json:"data_bucket,omitempty"`
}

func (o UpdateTrackerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrackerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTrackerRequestBody", string(data)}, " ")
}

type UpdateTrackerRequestBodyTrackerType struct {
	value string
}

type UpdateTrackerRequestBodyTrackerTypeEnum struct {
	SYSTEM UpdateTrackerRequestBodyTrackerType
	DATA   UpdateTrackerRequestBodyTrackerType
}

func GetUpdateTrackerRequestBodyTrackerTypeEnum() UpdateTrackerRequestBodyTrackerTypeEnum {
	return UpdateTrackerRequestBodyTrackerTypeEnum{
		SYSTEM: UpdateTrackerRequestBodyTrackerType{
			value: "system",
		},
		DATA: UpdateTrackerRequestBodyTrackerType{
			value: "data",
		},
	}
}

func (c UpdateTrackerRequestBodyTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrackerRequestBodyTrackerType) UnmarshalJSON(b []byte) error {
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

type UpdateTrackerRequestBodyStatus struct {
	value string
}

type UpdateTrackerRequestBodyStatusEnum struct {
	ENABLED  UpdateTrackerRequestBodyStatus
	DISABLED UpdateTrackerRequestBodyStatus
}

func GetUpdateTrackerRequestBodyStatusEnum() UpdateTrackerRequestBodyStatusEnum {
	return UpdateTrackerRequestBodyStatusEnum{
		ENABLED: UpdateTrackerRequestBodyStatus{
			value: "enabled",
		},
		DISABLED: UpdateTrackerRequestBodyStatus{
			value: "disabled",
		},
	}
}

func (c UpdateTrackerRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrackerRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
