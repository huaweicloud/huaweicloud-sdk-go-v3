package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ExportServerLogsResponse Response Object
type ExportServerLogsResponse struct {

	// export task id
	ExportId *string `json:"export_id,omitempty"`

	// 导出进度
	Status *ExportServerLogsResponseStatus `json:"status,omitempty"`

	// 导出进度百分比，范围0-100，如\"10%\"
	Percentage *string `json:"percentage,omitempty"`

	// 导出开始时间，时间戳格式为ISO 8601，例如：2025-03-20T02:48:06+00:00
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	ContentDisposition *string `json:"Content-Disposition,omitempty"`

	ContentTransferEncoding *string `json:"Content-Transfer-Encoding,omitempty"`

	ContentType    *string `json:"Content-Type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportServerLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportServerLogsResponse struct{}"
	}

	return strings.Join([]string{"ExportServerLogsResponse", string(data)}, " ")
}

type ExportServerLogsResponseStatus struct {
	value string
}

type ExportServerLogsResponseStatusEnum struct {
	PROCESSING ExportServerLogsResponseStatus
	COMPLETED  ExportServerLogsResponseStatus
}

func GetExportServerLogsResponseStatusEnum() ExportServerLogsResponseStatusEnum {
	return ExportServerLogsResponseStatusEnum{
		PROCESSING: ExportServerLogsResponseStatus{
			value: "processing",
		},
		COMPLETED: ExportServerLogsResponseStatus{
			value: "completed",
		},
	}
}

func (c ExportServerLogsResponseStatus) Value() string {
	return c.value
}

func (c ExportServerLogsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportServerLogsResponseStatus) UnmarshalJSON(b []byte) error {
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
