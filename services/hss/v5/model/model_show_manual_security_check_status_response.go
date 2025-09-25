package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowManualSecurityCheckStatusResponse Response Object
type ShowManualSecurityCheckStatusResponse struct {

	// 体检状态
	ScanStatus *ShowManualSecurityCheckStatusResponseScanStatus `json:"scan_status,omitempty"`

	// 本次体检服务器数量
	TotalHostNum *int32 `json:"total_host_num,omitempty"`

	// 体检完成的服务器数量
	ScanedHostNum *int32 `json:"scaned_host_num,omitempty"`

	// 体检内容数量
	ContentNum *int32 `json:"content_num,omitempty"`

	// 体检进度百分比
	ScanProgress   *int32 `json:"scan_progress,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowManualSecurityCheckStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManualSecurityCheckStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowManualSecurityCheckStatusResponse", string(data)}, " ")
}

type ShowManualSecurityCheckStatusResponseScanStatus struct {
	value string
}

type ShowManualSecurityCheckStatusResponseScanStatusEnum struct {
	SCANNING ShowManualSecurityCheckStatusResponseScanStatus
	SCANED   ShowManualSecurityCheckStatusResponseScanStatus
	FAILED   ShowManualSecurityCheckStatusResponseScanStatus
}

func GetShowManualSecurityCheckStatusResponseScanStatusEnum() ShowManualSecurityCheckStatusResponseScanStatusEnum {
	return ShowManualSecurityCheckStatusResponseScanStatusEnum{
		SCANNING: ShowManualSecurityCheckStatusResponseScanStatus{
			value: "scanning",
		},
		SCANED: ShowManualSecurityCheckStatusResponseScanStatus{
			value: "scaned",
		},
		FAILED: ShowManualSecurityCheckStatusResponseScanStatus{
			value: "failed",
		},
	}
}

func (c ShowManualSecurityCheckStatusResponseScanStatus) Value() string {
	return c.value
}

func (c ShowManualSecurityCheckStatusResponseScanStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowManualSecurityCheckStatusResponseScanStatus) UnmarshalJSON(b []byte) error {
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
