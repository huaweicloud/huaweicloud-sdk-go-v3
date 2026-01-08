package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmTemplatesRequest Request Object
type ListAlarmTemplatesRequest struct {

	// **参数解释**： 分页偏移量 **约束限制**： 不涉及 **取值范围**： 整数，[0,10000] **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小 **约束限制**： 不涉及 **取值范围**： 整数，[1,100] **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 资源维度名称。 **约束限制**： 不涉及。 **取值范围**： 多维度用\",\"分隔，只能包含0-9、a-z、A-Z、_、-、#、/、(、），每个维度的最大长度为32。字符串总长度最小为1，最大为131。 **默认取值**： 不涉及。
	DimName *string `json:"dim_name,omitempty"`

	// **参数解释**： 模板类型。 **约束限制**： 不涉及。 **取值范围**： 枚举值。 - system：默认指标模板。 - custom： 自定义指标模板。    - system_event： 默认事件模板。 - custom_event： 自定义事件模板。    - system_custom_event： 全部事件模板。     **默认取值**： 不传返回全部指标模板。
	TemplateType *ListAlarmTemplatesRequestTemplateType `json:"template_type,omitempty"`

	// 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]，支持模糊匹配
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释**： 产品层级跨维规则创建时需要指明的规则产品名称，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,128]个字符。          **默认取值**： 不涉及。
	ProductName *string `json:"product_name,omitempty"`
}

func (o ListAlarmTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplatesRequest", string(data)}, " ")
}

type ListAlarmTemplatesRequestTemplateType struct {
	value string
}

type ListAlarmTemplatesRequestTemplateTypeEnum struct {
	SYSTEM              ListAlarmTemplatesRequestTemplateType
	CUSTOM              ListAlarmTemplatesRequestTemplateType
	SYSTEM_EVENT        ListAlarmTemplatesRequestTemplateType
	CUSTOM_EVENT        ListAlarmTemplatesRequestTemplateType
	SYSTEM_CUSTOM_EVENT ListAlarmTemplatesRequestTemplateType
}

func GetListAlarmTemplatesRequestTemplateTypeEnum() ListAlarmTemplatesRequestTemplateTypeEnum {
	return ListAlarmTemplatesRequestTemplateTypeEnum{
		SYSTEM: ListAlarmTemplatesRequestTemplateType{
			value: "system",
		},
		CUSTOM: ListAlarmTemplatesRequestTemplateType{
			value: "custom",
		},
		SYSTEM_EVENT: ListAlarmTemplatesRequestTemplateType{
			value: "system_event",
		},
		CUSTOM_EVENT: ListAlarmTemplatesRequestTemplateType{
			value: "custom_event",
		},
		SYSTEM_CUSTOM_EVENT: ListAlarmTemplatesRequestTemplateType{
			value: "system_custom_event",
		},
	}
}

func (c ListAlarmTemplatesRequestTemplateType) Value() string {
	return c.value
}

func (c ListAlarmTemplatesRequestTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmTemplatesRequestTemplateType) UnmarshalJSON(b []byte) error {
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
