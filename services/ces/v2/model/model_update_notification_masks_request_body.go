package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateNotificationMasksRequestBody 通知屏蔽请求体
type UpdateNotificationMasksRequestBody struct {

	// **参数解释**： 屏蔽规则名称。    **约束限制**： 不涉及。 **取值范围**： 只能为字母、数字、汉字、-、_，长度为[1,64]个字符。      **默认取值**： 不涉及。
	MaskName string `json:"mask_name"`

	// **参数解释**： 关联ID，为告警规则ID或者告警策略ID    **约束限制**： 包含的关联ID数量为[1,100]个。 relation_type为RESOURCE_POLICY_NOTIFICATION时填屏蔽的告警策略ID。
	RelationIds *[]string `json:"relation_ids,omitempty"`

	RelationType *RelationType `json:"relation_type,omitempty"`

	// **参数解释** 关联指标名称列表。 **约束限制**： relation_type为RESOURCE可选填，不填视为对资源所有指标进行告警屏蔽。包含的指标名称最多为50个，最少为0个。
	MetricNames *[]string `json:"metric_names,omitempty"`

	// **参数解释** 按云产品维度屏蔽时的指标信息。 **约束限制**：| 包含的指标信息最多为50个，最少为0个。
	ProductMetrics *[]ProductMetric `json:"product_metrics,omitempty"`

	// **参数解释** 资源层级。 **约束限制**： 不涉及。 **取值范围**： 枚举值，取值为dimension、product - dimension: 子维度 - product: 云产品 **默认取值**： 不涉及。
	ResourceLevel *UpdateNotificationMasksRequestBodyResourceLevel `json:"resource_level,omitempty"`

	// **参数解释** 资源为云产品时的云产品名称。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,128]个字符。 **默认取值**： 不涉及。
	ProductName *string `json:"product_name,omitempty"`

	// **参数解释**： 关联资源。           **约束限制**： 数组长度为[1,100]
	Resources []Resource `json:"resources"`

	MaskType *MaskType `json:"mask_type"`

	// **参数解释**： 屏蔽起始日期。           **约束限制**： 不涉及。 **取值范围**： 字符长度为10，格式为：yyyy-MM-dd           **默认取值**： 不涉及。
	StartDate *string `json:"start_date,omitempty"`

	// **参数解释**： 屏蔽起始时间。          **约束限制**： 不涉及。 **取值范围**： 字符长度为8，格式为：HH:mm:ss         **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 屏蔽截止日期。           **约束限制**： 不涉及。 **取值范围**： 字符长度为10，格式为：yyyy-MM-dd           **默认取值**： 不涉及。
	EndDate *string `json:"end_date,omitempty"`

	// **参数解释**： 屏蔽截止时间。          **约束限制**： 不涉及。 **取值范围**： 字符长度为8，格式为：HH:mm:ss         **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **约束限制**： 不涉及。 **取值范围**： 长度为[1,16]个字符。           **默认取值**： 不涉及。
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`
}

func (o UpdateNotificationMasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationMasksRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNotificationMasksRequestBody", string(data)}, " ")
}

type UpdateNotificationMasksRequestBodyResourceLevel struct {
	value string
}

type UpdateNotificationMasksRequestBodyResourceLevelEnum struct {
	DIMENSION UpdateNotificationMasksRequestBodyResourceLevel
	PRODUCT   UpdateNotificationMasksRequestBodyResourceLevel
}

func GetUpdateNotificationMasksRequestBodyResourceLevelEnum() UpdateNotificationMasksRequestBodyResourceLevelEnum {
	return UpdateNotificationMasksRequestBodyResourceLevelEnum{
		DIMENSION: UpdateNotificationMasksRequestBodyResourceLevel{
			value: "dimension",
		},
		PRODUCT: UpdateNotificationMasksRequestBodyResourceLevel{
			value: "product",
		},
	}
}

func (c UpdateNotificationMasksRequestBodyResourceLevel) Value() string {
	return c.value
}

func (c UpdateNotificationMasksRequestBodyResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationMasksRequestBodyResourceLevel) UnmarshalJSON(b []byte) error {
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
