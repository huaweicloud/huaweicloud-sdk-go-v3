package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowLogsExportStatusResponse Response Object
type ShowLogsExportStatusResponse struct {

	// export task id
	ExportId *string `json:"export_id,omitempty"`

	// 导出进度
	Status *ShowLogsExportStatusResponseStatus `json:"status,omitempty"`

	// 导出进度百分比，范围0-100，如\"10%\"
	Percentage *string `json:"percentage,omitempty"`

	// 导出开始时间，时间戳格式为ISO 8601，例如：2025-03-20T02:48:06+00:00
	CreatedAt      *sdktime.SdkTime `json:"created_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowLogsExportStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogsExportStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowLogsExportStatusResponse", string(data)}, " ")
}

type ShowLogsExportStatusResponseStatus struct {
	value string
}

type ShowLogsExportStatusResponseStatusEnum struct {
	PROCESSING ShowLogsExportStatusResponseStatus
	COMPLETED  ShowLogsExportStatusResponseStatus
}

func GetShowLogsExportStatusResponseStatusEnum() ShowLogsExportStatusResponseStatusEnum {
	return ShowLogsExportStatusResponseStatusEnum{
		PROCESSING: ShowLogsExportStatusResponseStatus{
			value: "processing",
		},
		COMPLETED: ShowLogsExportStatusResponseStatus{
			value: "completed",
		},
	}
}

func (c ShowLogsExportStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowLogsExportStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLogsExportStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
