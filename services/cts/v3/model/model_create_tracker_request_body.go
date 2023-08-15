package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTrackerRequestBody struct {

	// 标识追踪器类型。 目前支持系统追踪器类型有管理类追踪器(system)和数据类追踪器(data)。 数据类追踪器和管理类追踪器共同参数有：is_lts_enabled, obs_info; 管理类追踪器参数：is_support_trace_files_encryption, kms_id, is_support_validate, is_support_validate; 数据类追踪器参数：tracker_name, data_bucket。
	TrackerType CreateTrackerRequestBodyTrackerType `json:"tracker_type"`

	// 标识追踪器名称。 当\"tracker_type\"参数值为\"system\"时该参数为默认值\"system\"。 当\"tracker_type\"参数值为\"data\"时该参数需要指定追踪器名称\"。
	TrackerName string `json:"tracker_name"`

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

func (o CreateTrackerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTrackerRequestBody", string(data)}, " ")
}

type CreateTrackerRequestBodyTrackerType struct {
	value string
}

type CreateTrackerRequestBodyTrackerTypeEnum struct {
	SYSTEM CreateTrackerRequestBodyTrackerType
	DATA   CreateTrackerRequestBodyTrackerType
}

func GetCreateTrackerRequestBodyTrackerTypeEnum() CreateTrackerRequestBodyTrackerTypeEnum {
	return CreateTrackerRequestBodyTrackerTypeEnum{
		SYSTEM: CreateTrackerRequestBodyTrackerType{
			value: "system",
		},
		DATA: CreateTrackerRequestBodyTrackerType{
			value: "data",
		},
	}
}

func (c CreateTrackerRequestBodyTrackerType) Value() string {
	return c.value
}

func (c CreateTrackerRequestBodyTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerRequestBodyTrackerType) UnmarshalJSON(b []byte) error {
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
