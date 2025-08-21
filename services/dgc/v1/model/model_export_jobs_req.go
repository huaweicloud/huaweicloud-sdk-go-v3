package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportJobsReq struct {
	JobList *[]string `json:"jobList,omitempty"`

	// 是否导出作业依赖的脚本和资源，取值为true或者false，不传该值时，默认为true。
	ExportDepend *bool `json:"exportDepend,omitempty"`

	// 当导出到obs时，需要指定obs路径，样例：obs://test_bucket/job_nodes/
	ObsPath *string `json:"obsPath,omitempty"`

	// 导出作业的状态，取值如下： - DEVELOP: 开发态，默认是开发态 - SUBMIT: 提交态
	ExportStatus *ExportJobsReqExportStatus `json:"exportStatus,omitempty"`
}

func (o ExportJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobsReq struct{}"
	}

	return strings.Join([]string{"ExportJobsReq", string(data)}, " ")
}

type ExportJobsReqExportStatus struct {
	value string
}

type ExportJobsReqExportStatusEnum struct {
	SUBMIT  ExportJobsReqExportStatus
	DEVELOP ExportJobsReqExportStatus
}

func GetExportJobsReqExportStatusEnum() ExportJobsReqExportStatusEnum {
	return ExportJobsReqExportStatusEnum{
		SUBMIT: ExportJobsReqExportStatus{
			value: "SUBMIT",
		},
		DEVELOP: ExportJobsReqExportStatus{
			value: "DEVELOP",
		},
	}
}

func (c ExportJobsReqExportStatus) Value() string {
	return c.value
}

func (c ExportJobsReqExportStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqExportStatus) UnmarshalJSON(b []byte) error {
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
