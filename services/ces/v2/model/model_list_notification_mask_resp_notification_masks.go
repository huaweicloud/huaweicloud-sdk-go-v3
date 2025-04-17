package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListNotificationMaskRespNotificationMasks struct {

	// 屏蔽规则ID
	NotificationMaskId string `json:"notification_mask_id"`

	// 屏蔽规则名称，只能为字母、数字、汉字、-、_，最大长度为64
	MaskName *string `json:"mask_name,omitempty"`

	RelationType *RelationType `json:"relation_type"`

	// 关联编号
	RelationId *string `json:"relation_id,omitempty"`

	ResourceType *MaskResourceType `json:"resource_type,omitempty"`

	// 关联指标名称，relation_type为RESOURCE时存在该字段
	MetricNames *[]string `json:"metric_names,omitempty"`

	// 按云产品维度屏蔽时的指标信息
	ProductMetrics *[]ProductMetric `json:"product_metrics,omitempty"`

	// dimension: 子维度,product: 云产品
	ResourceLevel *ListNotificationMaskRespNotificationMasksResourceLevel `json:"resource_level,omitempty"`

	// 资源为云产品时的云产品名称
	ProductName *string `json:"product_name,omitempty"`

	// 关联资源类型，relation_type为RESOURCE时存在该字段,只需要查询出资源的namespace+维度名即可
	Resources *[]ResourceCategory `json:"resources,omitempty"`

	MaskStatus *MaskStatus `json:"mask_status"`

	MaskType *MaskType `json:"mask_type"`

	// 告警屏蔽的创建时间，UNIX时间戳，单位毫秒。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 告警屏蔽的更新时间，UNIX时间戳，单位毫秒。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 屏蔽起始日期，yyyy-MM-dd。
	StartDate *string `json:"start_date,omitempty"`

	// 屏蔽起始时间，HH:mm:ss。
	StartTime *string `json:"start_time,omitempty"`

	// 屏蔽截止日期，yyyy-MM-dd。
	EndDate *string `json:"end_date,omitempty"`

	// 屏蔽截止时间，HH:mm:ss。
	EndTime *string `json:"end_time,omitempty"`

	// 告警策略列表。
	Policies *[]PoliciesInListResp `json:"policies,omitempty"`
}

func (o ListNotificationMaskRespNotificationMasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMaskRespNotificationMasks struct{}"
	}

	return strings.Join([]string{"ListNotificationMaskRespNotificationMasks", string(data)}, " ")
}

type ListNotificationMaskRespNotificationMasksResourceLevel struct {
	value string
}

type ListNotificationMaskRespNotificationMasksResourceLevelEnum struct {
	DIMENSION ListNotificationMaskRespNotificationMasksResourceLevel
	PRODUCT   ListNotificationMaskRespNotificationMasksResourceLevel
}

func GetListNotificationMaskRespNotificationMasksResourceLevelEnum() ListNotificationMaskRespNotificationMasksResourceLevelEnum {
	return ListNotificationMaskRespNotificationMasksResourceLevelEnum{
		DIMENSION: ListNotificationMaskRespNotificationMasksResourceLevel{
			value: "dimension",
		},
		PRODUCT: ListNotificationMaskRespNotificationMasksResourceLevel{
			value: "product",
		},
	}
}

func (c ListNotificationMaskRespNotificationMasksResourceLevel) Value() string {
	return c.value
}

func (c ListNotificationMaskRespNotificationMasksResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationMaskRespNotificationMasksResourceLevel) UnmarshalJSON(b []byte) error {
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
