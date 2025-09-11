package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateResourceGroupRequestBody struct {

	// **参数解释** 资源分组的名称 **约束限制** 不涉及 **取值范围** 只能为字母、数字、汉字、-或_，长度为[1,128]个字符 **默认取值** 不涉及
	GroupName string `json:"group_name"`

	// **参数解释** 资源分组归属企业项目ID **约束限制** 不涉及 **取值范围** 由数字、字母和-组成，或者为0（默认企业项目ID）。 **默认取值** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 资源分组添加资源方式 **约束限制** 不涉及 **取值范围** 取值只能为EPS（同步企业项目），TAG（标签动态匹配），NAME（实例名称），COMB（组合匹配），不传为手动添加。 **默认取值** 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释** 标签动态匹配时的关联标签。 **约束限制** type为TAG时必传，不超过50个标签。
	Tags *[]ResourceGroupTagRelation `json:"tags,omitempty"`

	// **参数解释** 该资源分组内包含的资源来源的企业项目ID。 **约束限制** type为EPS时必传，不超过50个企业项目ID。
	AssociationEpIds *[]string `json:"association_ep_ids,omitempty"`

	// **参数解释** 云服务名称,格式为\"dcs,ecs\",支持的云服务providers请参考《配置审计API参考》中的\"支持的服务和资源类型\"章节 **约束限制** 不涉及 **取值范围** 长度为[0,512]个字符 **默认取值** 不涉及
	Providers *string `json:"providers,omitempty"`

	// **参数解释** 匹配企业项目或匹配标签参数。 **约束限制** 不超过50个条件。
	EnterpriseProjectIdAndTags *[]EnterpriseProjectIdAndTags `json:"enterprise_project_id_and_tags,omitempty"`

	// **参数解释** 手动创建时的资源详情。 **约束限制** 不超过1000个资源。
	Resources *[]Resource `json:"resources,omitempty"`

	// **参数解释** 手动创建，选择资源层级为云产品时的资源详情。 **约束限制** 不超过50个资源。
	ProductResources *[]ProductResource `json:"product_resources,omitempty"`

	// **参数解释** 实例名称匹配参数。 **约束限制** type为NAME时必传，不超过50个实例。
	Instances *[]Instance `json:"instances,omitempty"`

	// **参数解释** 创建资源层级为云产品时的云产品的取值，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。多个云产品则用“;”隔开，如\"SERVICE.BMS,instance_id;SYS.ECS,instance_id\"。 **约束限制** 不涉及。 **取值范围** 长度[0,10240]个字符 **默认取值** 不涉及。
	ProductNames *string `json:"product_names,omitempty"`

	// **参数解释** 资源层级，资源生效范围。选择云产品，则云产品及其子层级均可进入该资源分组，选择子维度，则只生效具体的子维度。 **约束限制** 不涉及。 **取值范围** - product: 云产品 - dimension: 子维度 **默认取值** 不涉及。
	ResourceLevel *CreateResourceGroupRequestBodyResourceLevel `json:"resource_level,omitempty"`

	CombRelation *CombRelation `json:"comb_relation,omitempty"`
}

func (o CreateResourceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupRequestBody", string(data)}, " ")
}

type CreateResourceGroupRequestBodyResourceLevel struct {
	value string
}

type CreateResourceGroupRequestBodyResourceLevelEnum struct {
	PRODUCT   CreateResourceGroupRequestBodyResourceLevel
	DIMENSION CreateResourceGroupRequestBodyResourceLevel
}

func GetCreateResourceGroupRequestBodyResourceLevelEnum() CreateResourceGroupRequestBodyResourceLevelEnum {
	return CreateResourceGroupRequestBodyResourceLevelEnum{
		PRODUCT: CreateResourceGroupRequestBodyResourceLevel{
			value: "product",
		},
		DIMENSION: CreateResourceGroupRequestBodyResourceLevel{
			value: "dimension",
		},
	}
}

func (c CreateResourceGroupRequestBodyResourceLevel) Value() string {
	return c.value
}

func (c CreateResourceGroupRequestBodyResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateResourceGroupRequestBodyResourceLevel) UnmarshalJSON(b []byte) error {
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
