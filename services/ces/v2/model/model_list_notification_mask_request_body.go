package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListNotificationMaskRequestBody struct {
	RelationType *ListRelationType `json:"relation_type"`

	// 关联编号（目前是告警规则ID）
	RelationIds []string `json:"relation_ids"`

	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。
	MetricName *string `json:"metric_name,omitempty"`

	// dimension: 子维度,product: 云产品
	ResourceLevel *ListNotificationMaskRequestBodyResourceLevel `json:"resource_level,omitempty"`

	// 屏蔽规则ID,可选
	MaskId *string `json:"mask_id,omitempty"`

	// 屏蔽规则名称,可选，只能为字母、数字、汉字、-、_，最大长度为64
	MaskName *string `json:"mask_name,omitempty"`

	// 屏蔽状态,可选。MASK_EFFECTIVE：已生效，MASK_INEFFECTIVE：未生效。
	MaskStatus *ListNotificationMaskRequestBodyMaskStatus `json:"mask_status,omitempty"`

	// 资源维度值,提供一个维度的资源ID即可,可选
	ResourceId *string `json:"resource_id,omitempty"`

	// 服务的命名空间，查询各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace *string `json:"namespace,omitempty"`

	// 资源的维度信息
	Dimensions *[]ResourceDimension `json:"dimensions,omitempty"`
}

func (o ListNotificationMaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMaskRequestBody struct{}"
	}

	return strings.Join([]string{"ListNotificationMaskRequestBody", string(data)}, " ")
}

type ListNotificationMaskRequestBodyResourceLevel struct {
	value string
}

type ListNotificationMaskRequestBodyResourceLevelEnum struct {
	DIMENSION ListNotificationMaskRequestBodyResourceLevel
	PRODUCT   ListNotificationMaskRequestBodyResourceLevel
}

func GetListNotificationMaskRequestBodyResourceLevelEnum() ListNotificationMaskRequestBodyResourceLevelEnum {
	return ListNotificationMaskRequestBodyResourceLevelEnum{
		DIMENSION: ListNotificationMaskRequestBodyResourceLevel{
			value: "dimension",
		},
		PRODUCT: ListNotificationMaskRequestBodyResourceLevel{
			value: "product",
		},
	}
}

func (c ListNotificationMaskRequestBodyResourceLevel) Value() string {
	return c.value
}

func (c ListNotificationMaskRequestBodyResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationMaskRequestBodyResourceLevel) UnmarshalJSON(b []byte) error {
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

type ListNotificationMaskRequestBodyMaskStatus struct {
	value string
}

type ListNotificationMaskRequestBodyMaskStatusEnum struct {
	MASK_EFFECTIVE   ListNotificationMaskRequestBodyMaskStatus
	MASK_INEFFECTIVE ListNotificationMaskRequestBodyMaskStatus
}

func GetListNotificationMaskRequestBodyMaskStatusEnum() ListNotificationMaskRequestBodyMaskStatusEnum {
	return ListNotificationMaskRequestBodyMaskStatusEnum{
		MASK_EFFECTIVE: ListNotificationMaskRequestBodyMaskStatus{
			value: "MASK_EFFECTIVE",
		},
		MASK_INEFFECTIVE: ListNotificationMaskRequestBodyMaskStatus{
			value: "MASK_INEFFECTIVE",
		},
	}
}

func (c ListNotificationMaskRequestBodyMaskStatus) Value() string {
	return c.value
}

func (c ListNotificationMaskRequestBodyMaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationMaskRequestBodyMaskStatus) UnmarshalJSON(b []byte) error {
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
