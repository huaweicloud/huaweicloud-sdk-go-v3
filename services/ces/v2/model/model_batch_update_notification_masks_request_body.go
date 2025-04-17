package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdateNotificationMasksRequestBody 通知屏蔽信请求体
type BatchUpdateNotificationMasksRequestBody struct {

	// 屏蔽规则名称，只能为字母、数字、汉字、-、_，最大长度为64
	MaskName *string `json:"mask_name,omitempty"`

	RelationType *RelationType `json:"relation_type"`

	// 关联编号，relation_type为ALARM_RULE时填屏蔽的告警规则ID；relation_type为RESOURCE_POLICY_NOTIFICATION、RESOURCE_POLICY_ALARM时填屏蔽的告警策略ID；
	RelationIds []string `json:"relation_ids"`

	// 关联资源，relation_type为RESOURCE、RESOURCE_POLICY_NOTIFICATION、RESOURCE_POLICY_ALARM时填屏蔽的资源信息；
	Resources *[]Resource `json:"resources,omitempty"`

	// 关联指标名称，relation_type为RESOURCE可选填，不填视为对资源所有指标进行告警屏蔽
	MetricNames *[]string `json:"metric_names,omitempty"`

	// 按云产品维度屏蔽时的指标信息
	ProductMetrics *[]ProductMetric `json:"product_metrics,omitempty"`

	// dimension: 子维度,product: 云产品
	ResourceLevel *BatchUpdateNotificationMasksRequestBodyResourceLevel `json:"resource_level,omitempty"`

	// 资源为云产品时的云产品名称
	ProductName *string `json:"product_name,omitempty"`

	MaskType *MaskType `json:"mask_type"`

	// 屏蔽起始日期，yyyy-MM-dd。
	StartDate *string `json:"start_date,omitempty"`

	// 屏蔽起始时间，HH:mm:ss。
	StartTime *string `json:"start_time,omitempty"`

	// 屏蔽截止日期，yyyy-MM-dd。
	EndDate *string `json:"end_date,omitempty"`

	// 屏蔽截止时间，HH:mm:ss。
	EndTime *string `json:"end_time,omitempty"`
}

func (o BatchUpdateNotificationMasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNotificationMasksRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateNotificationMasksRequestBody", string(data)}, " ")
}

type BatchUpdateNotificationMasksRequestBodyResourceLevel struct {
	value string
}

type BatchUpdateNotificationMasksRequestBodyResourceLevelEnum struct {
	DIMENSION BatchUpdateNotificationMasksRequestBodyResourceLevel
	PRODUCT   BatchUpdateNotificationMasksRequestBodyResourceLevel
}

func GetBatchUpdateNotificationMasksRequestBodyResourceLevelEnum() BatchUpdateNotificationMasksRequestBodyResourceLevelEnum {
	return BatchUpdateNotificationMasksRequestBodyResourceLevelEnum{
		DIMENSION: BatchUpdateNotificationMasksRequestBodyResourceLevel{
			value: "dimension",
		},
		PRODUCT: BatchUpdateNotificationMasksRequestBodyResourceLevel{
			value: "product",
		},
	}
}

func (c BatchUpdateNotificationMasksRequestBodyResourceLevel) Value() string {
	return c.value
}

func (c BatchUpdateNotificationMasksRequestBodyResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateNotificationMasksRequestBodyResourceLevel) UnmarshalJSON(b []byte) error {
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
