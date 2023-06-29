package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckAssetJobStatusResponse Response Object
type CheckAssetJobStatusResponse struct {

	// 作业ID
	Id *string `json:"id,omitempty"`

	// 作业类型
	Type *CheckAssetJobStatusResponseType `json:"type,omitempty"`

	// 作业状态 - RUNNING : 作业正在执行 - SUCCEEDED : 作业执行成功，对于导出作业，用户可以通过archive_id来下载资产包 - FAILED : 作业执行失败，通过reason字段查看具体错误原因
	Status *CheckAssetJobStatusResponseStatus `json:"status,omitempty"`

	// 导致作业失败的错误原因
	Reasons *[]AssetJobReason `json:"reasons,omitempty"`

	// 作业进度百分比
	ProgressPercent *float32 `json:"progress_percent,omitempty"`

	// 导出作业成功时，供下载的资产包ID
	ArchiveId *string `json:"archive_id,omitempty"`

	// 作业开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 作业结束时间
	EndTime        *string `json:"end_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckAssetJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAssetJobStatusResponse struct{}"
	}

	return strings.Join([]string{"CheckAssetJobStatusResponse", string(data)}, " ")
}

type CheckAssetJobStatusResponseType struct {
	value string
}

type CheckAssetJobStatusResponseTypeEnum struct {
	IMPORT_ASSET CheckAssetJobStatusResponseType
	EXPORT_ASSET CheckAssetJobStatusResponseType
}

func GetCheckAssetJobStatusResponseTypeEnum() CheckAssetJobStatusResponseTypeEnum {
	return CheckAssetJobStatusResponseTypeEnum{
		IMPORT_ASSET: CheckAssetJobStatusResponseType{
			value: "importAsset",
		},
		EXPORT_ASSET: CheckAssetJobStatusResponseType{
			value: "exportAsset",
		},
	}
}

func (c CheckAssetJobStatusResponseType) Value() string {
	return c.value
}

func (c CheckAssetJobStatusResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckAssetJobStatusResponseType) UnmarshalJSON(b []byte) error {
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

type CheckAssetJobStatusResponseStatus struct {
	value string
}

type CheckAssetJobStatusResponseStatusEnum struct {
	RUNNING   CheckAssetJobStatusResponseStatus
	SUCCEEDED CheckAssetJobStatusResponseStatus
	FAILED    CheckAssetJobStatusResponseStatus
}

func GetCheckAssetJobStatusResponseStatusEnum() CheckAssetJobStatusResponseStatusEnum {
	return CheckAssetJobStatusResponseStatusEnum{
		RUNNING: CheckAssetJobStatusResponseStatus{
			value: "RUNNING",
		},
		SUCCEEDED: CheckAssetJobStatusResponseStatus{
			value: "SUCCEEDED",
		},
		FAILED: CheckAssetJobStatusResponseStatus{
			value: "FAILED",
		},
	}
}

func (c CheckAssetJobStatusResponseStatus) Value() string {
	return c.value
}

func (c CheckAssetJobStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckAssetJobStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
