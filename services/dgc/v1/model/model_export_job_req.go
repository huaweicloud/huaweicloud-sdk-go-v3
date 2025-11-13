package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportJobReq struct {

	// 是否导出作业依赖的脚本和资源，取值为true或者false，不传该值时，默认为true。
	ExportDepend *bool `json:"exportDepend,omitempty"`

	// 当导出到obs时，需要指定obs路径，样例：obs://test_bucket/job_nodes/
	ObsPath *string `json:"obsPath,omitempty"`

	// 导出作业的状态，取值如下： - DEVELOP: 开发态，默认是开发态 - SUBMIT: 提交态
	ExportStatus *ExportJobReqExportStatus `json:"exportStatus,omitempty"`
}

func (o ExportJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobReq struct{}"
	}

	return strings.Join([]string{"ExportJobReq", string(data)}, " ")
}

type ExportJobReqExportStatus struct {
	value string
}

type ExportJobReqExportStatusEnum struct {
	SUBMIT  ExportJobReqExportStatus
	DEVELOP ExportJobReqExportStatus
}

func GetExportJobReqExportStatusEnum() ExportJobReqExportStatusEnum {
	return ExportJobReqExportStatusEnum{
		SUBMIT: ExportJobReqExportStatus{
			value: "SUBMIT",
		},
		DEVELOP: ExportJobReqExportStatus{
			value: "DEVELOP",
		},
	}
}

func (c ExportJobReqExportStatus) Value() string {
	return c.value
}

func (c ExportJobReqExportStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobReqExportStatus) UnmarshalJSON(b []byte) error {
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
