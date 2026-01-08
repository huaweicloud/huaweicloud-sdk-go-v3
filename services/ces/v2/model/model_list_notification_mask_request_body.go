package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListNotificationMaskRequestBody struct {
	RelationType *ListRelationType `json:"relation_type"`

	// **参数解释**： 告警规则ID列表，用于查询对应的告警通知屏蔽。 **约束限制**： 当relation_type为ALARM_RULE、RESOURCE_POLICY_NOTIFICATION时，应通过alarm_ids查询。当relation_type为RESOURCE、EVENT.SYS时，不支持使用alarm_ids查询，此时alarm_ids为空或不选，表示查询所有的RESOURCE、EVENT.SYS类型的告警屏蔽。 **取值范围**： 包含的告警规则ID数量最多不超过100个，最少为0个。 **默认取值**： 不涉及
	AlarmIds *[]string `json:"alarm_ids,omitempty"`

	// （已废弃，不推荐使用）关联编号（目前是告警规则ID）
	RelationIds []string `json:"relation_ids"`

	// **参数解释**： 资源的监控指标名称，各服务资源的指标名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-。字符长度最短为1，最大为96。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。         **默认取值**： 不涉及。
	MetricName *string `json:"metric_name,omitempty"`

	// **参数解释**： 资源层级。 **约束限制**： 不涉及。 **取值范围**： - product: 云产品   - dimension: 子维度 **默认取值**： 不涉及。
	ResourceLevel *ListNotificationMaskRequestBodyResourceLevel `json:"resource_level,omitempty"`

	// **参数解释**： 屏蔽规则ID,可选。 **约束限制**： 不涉及。 **取值范围**： 以nm开头，后跟[0,62]个英文或数字。 **默认取值**： 不涉及。
	MaskId *string `json:"mask_id,omitempty"`

	// **参数解释**： 屏蔽规则名称,可选。 **约束限制**： 不涉及。 **取值范围**： 只能为字母、数字、汉字、-、_，长度为[1,64]个字符。 **默认取值**： 不涉及。
	MaskName *string `json:"mask_name,omitempty"`

	// **参数解释**： 屏蔽状态,可选。 **约束限制**： 不涉及。 **取值范围**： - MASK_EFFECTIVE：已生效。 - MASK_INEFFECTIVE：未生效。 **默认取值**： 不涉及。
	MaskStatus *ListNotificationMaskRequestBodyMaskStatus `json:"mask_status,omitempty"`

	// **参数解释**： 资源维度值,提供一个维度的资源ID即可,可选 **约束限制**： 不涉及。 **取值范围**： 长度为[1,700]个字符。 **默认取值**： 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 资源的维度信息 **约束限制**： 包含的维度信息数量为[1,4]个。
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
