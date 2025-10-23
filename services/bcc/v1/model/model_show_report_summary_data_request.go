package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReportSummaryDataRequest Request Object
type ShowReportSummaryDataRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 报告ID
	ReportId string `json:"report_id"`

	// 数据项名称，取值范围：task-status-pie,task-type-pie,task-status-region-column,task-trend-curve,resource-protection-pie,resource-compliance-pie,resource-protection-region-column,resource-compliance-region-column,resource-trend-curve
	DataItemName ShowReportSummaryDataRequestDataItemName `json:"data_item_name"`
}

func (o ShowReportSummaryDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportSummaryDataRequest struct{}"
	}

	return strings.Join([]string{"ShowReportSummaryDataRequest", string(data)}, " ")
}

type ShowReportSummaryDataRequestDataItemName struct {
	value string
}

type ShowReportSummaryDataRequestDataItemNameEnum struct {
	TASK_STATUS_PIE                   ShowReportSummaryDataRequestDataItemName
	TASK_TYPE_PIE                     ShowReportSummaryDataRequestDataItemName
	TASK_STATUS_REGION_COLUMN         ShowReportSummaryDataRequestDataItemName
	TASK_TREND_CURVE                  ShowReportSummaryDataRequestDataItemName
	RESOURCE_PROTECTION_PIE           ShowReportSummaryDataRequestDataItemName
	RESOURCE_COMPLIANCE_PIE           ShowReportSummaryDataRequestDataItemName
	RESOURCE_PROTECTION_REGION_COLUMN ShowReportSummaryDataRequestDataItemName
	RESOURCE_COMPLIANCE_REGION_COLUMN ShowReportSummaryDataRequestDataItemName
	RESOURCE_TREND_CURVE              ShowReportSummaryDataRequestDataItemName
}

func GetShowReportSummaryDataRequestDataItemNameEnum() ShowReportSummaryDataRequestDataItemNameEnum {
	return ShowReportSummaryDataRequestDataItemNameEnum{
		TASK_STATUS_PIE: ShowReportSummaryDataRequestDataItemName{
			value: "task-status-pie",
		},
		TASK_TYPE_PIE: ShowReportSummaryDataRequestDataItemName{
			value: "task-type-pie",
		},
		TASK_STATUS_REGION_COLUMN: ShowReportSummaryDataRequestDataItemName{
			value: "task-status-region-column",
		},
		TASK_TREND_CURVE: ShowReportSummaryDataRequestDataItemName{
			value: "task-trend-curve",
		},
		RESOURCE_PROTECTION_PIE: ShowReportSummaryDataRequestDataItemName{
			value: "resource-protection-pie",
		},
		RESOURCE_COMPLIANCE_PIE: ShowReportSummaryDataRequestDataItemName{
			value: "resource-compliance-pie",
		},
		RESOURCE_PROTECTION_REGION_COLUMN: ShowReportSummaryDataRequestDataItemName{
			value: "resource-protection-region-column",
		},
		RESOURCE_COMPLIANCE_REGION_COLUMN: ShowReportSummaryDataRequestDataItemName{
			value: "resource-compliance-region-column",
		},
		RESOURCE_TREND_CURVE: ShowReportSummaryDataRequestDataItemName{
			value: "resource-trend-curve",
		},
	}
}

func (c ShowReportSummaryDataRequestDataItemName) Value() string {
	return c.value
}

func (c ShowReportSummaryDataRequestDataItemName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReportSummaryDataRequestDataItemName) UnmarshalJSON(b []byte) error {
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
