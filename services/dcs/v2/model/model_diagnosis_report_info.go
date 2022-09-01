package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DiagnosisReportInfo struct {

	// 诊断报告ID
	ReportId string `json:"report_id" xml:"report_id"`

	// 诊断任务状态
	Status DiagnosisReportInfoStatus `json:"status" xml:"status"`

	// 诊断时间段的开始时间。格式为：2017-03-31T12:24:46.297Z
	BeginTime string `json:"begin_time" xml:"begin_time"`

	// 诊断时间段的结束时间。格式为：2017-03-31T12:24:46.297Z
	EndTime string `json:"end_time" xml:"end_time"`

	// 诊断报告创建时间
	CreatedAt string `json:"created_at" xml:"created_at"`

	// 参与诊断的节点个数
	NodeNum int32 `json:"node_num" xml:"node_num"`

	// 诊断结果为异常的诊断项总数
	AbnormalItemSum int32 `json:"abnormal_item_sum" xml:"abnormal_item_sum"`

	// 诊断失败的诊断项总数
	FailedItemSum int32 `json:"failed_item_sum" xml:"failed_item_sum"`
}

func (o DiagnosisReportInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisReportInfo struct{}"
	}

	return strings.Join([]string{"DiagnosisReportInfo", string(data)}, " ")
}

type DiagnosisReportInfoStatus struct {
	value string
}

type DiagnosisReportInfoStatusEnum struct {
	DIAGNOSING DiagnosisReportInfoStatus
	FINISHED   DiagnosisReportInfoStatus
}

func GetDiagnosisReportInfoStatusEnum() DiagnosisReportInfoStatusEnum {
	return DiagnosisReportInfoStatusEnum{
		DIAGNOSING: DiagnosisReportInfoStatus{
			value: "diagnosing",
		},
		FINISHED: DiagnosisReportInfoStatus{
			value: "finished",
		},
	}
}

func (c DiagnosisReportInfoStatus) Value() string {
	return c.value
}

func (c DiagnosisReportInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisReportInfoStatus) UnmarshalJSON(b []byte) error {
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
