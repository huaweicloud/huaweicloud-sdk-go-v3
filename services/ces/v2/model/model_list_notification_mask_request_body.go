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

	// 屏蔽规则ID,可选
	MaskId *string `json:"mask_id,omitempty"`

	// 屏蔽规则名称,可选，只能为字母、数字、汉字、-、_，最大长度为64
	MaskName *string `json:"mask_name,omitempty"`

	// 屏蔽状态,可选。MASK_EFFECTIVE：已生效，MASK_INEFFECTIVE：未生效。
	MaskStatus *ListNotificationMaskRequestBodyMaskStatus `json:"mask_status,omitempty"`

	// 资源维度值,提供一个维度的资源ID即可,可选
	ResourceId *string `json:"resource_id,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
	Namespace *string `json:"namespace,omitempty"`

	// 资源的维度信息
	Dimensions *[]Dimension2 `json:"dimensions,omitempty"`
}

func (o ListNotificationMaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMaskRequestBody struct{}"
	}

	return strings.Join([]string{"ListNotificationMaskRequestBody", string(data)}, " ")
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
