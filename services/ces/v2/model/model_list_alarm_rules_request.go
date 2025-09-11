package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmRulesRequest Request Object
type ListAlarmRulesRequest struct {

	// **参数解释**： 告警规则ID。 **约束限制**： 不涉及。 **取值范围**： 以al开头，后跟22位的字母或数字。          **默认取值**： 不涉及。
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警名称。 **约束限制**： 不涉及。 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度为[1,128]个字符。          **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度在 0 到 32个字符之间。           **默认取值**： 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 告警资源ID。 **约束限制**： 不涉及。 **取值范围**： 多维度情况按字母升序排列并使用逗号分隔。 长度为[0,700]个字符。        **默认取值**： 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 企业项目ID。 **约束限制**： 不涉及。 **取值范围**： 只能包含小写字母、数字、“-”、“_”，可以自定义企业项目ID，长度为36个字符。也可以为0（代表默认企业项目ID），all_granted_eps（代表所有企业项目ID）。           **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 产品层级跨维规则查询时支持产品名称查询，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,128]个字符。        **默认取值**： 不涉及。
	ProductName *string `json:"product_name,omitempty"`

	// **参数解释**： 产品层级跨维规则查询时支持规则所属类型查询，resource_level取值为product即为云产品类型，不填或者取值为dimension则为子维度类型。           **约束限制**： 不涉及。 **取值范围**： 枚举值。 - product：云产品。 - dimension：子维度。 **默认取值**： 不涉及。
	ResourceLevel *ListAlarmRulesRequestResourceLevel `json:"resource_level,omitempty"`

	// **参数解释**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： [0,10000]           **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小。 **约束限制**： 不涉及。 **取值范围**： [1,100]           **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesRequest", string(data)}, " ")
}

type ListAlarmRulesRequestResourceLevel struct {
	value string
}

type ListAlarmRulesRequestResourceLevelEnum struct {
	PRODUCT   ListAlarmRulesRequestResourceLevel
	DIMENSION ListAlarmRulesRequestResourceLevel
}

func GetListAlarmRulesRequestResourceLevelEnum() ListAlarmRulesRequestResourceLevelEnum {
	return ListAlarmRulesRequestResourceLevelEnum{
		PRODUCT: ListAlarmRulesRequestResourceLevel{
			value: "product",
		},
		DIMENSION: ListAlarmRulesRequestResourceLevel{
			value: "dimension",
		},
	}
}

func (c ListAlarmRulesRequestResourceLevel) Value() string {
	return c.value
}

func (c ListAlarmRulesRequestResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmRulesRequestResourceLevel) UnmarshalJSON(b []byte) error {
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
