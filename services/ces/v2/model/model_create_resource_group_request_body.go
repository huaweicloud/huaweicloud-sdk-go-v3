package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateResourceGroupRequestBody struct {

	// 资源分组的名称，只能为字母、数字、汉字、-、_，最大长度为128
	GroupName string `json:"group_name"`

	// 资源分组归属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源分组添加资源方式，取值只能为EPS（同步企业项目）,TAG（标签动态匹配）,NAME（实例名称）,不传为手动添加
	Type *string `json:"type,omitempty"`

	// 标签动态匹配时的关联标签,type为TAG时必传
	Tags *[]ResourceGroupTagRelation `json:"tags,omitempty"`

	// 该资源分组内包含的资源来源的企业项目ID，type为EPS时必传
	AssociationEpIds *[]string `json:"association_ep_ids,omitempty"`

	// 云服务名称,格式为\"dcs,ecs\",支持的云服务providers请参考https://support.huaweicloud.com/api-rms/rms_06_0100.html
	Providers *string `json:"providers,omitempty"`

	// 组合匹配参数
	EnterpriseProjectIdAndTags *[]EnterpriseProjectIdAndTags `json:"enterprise_project_id_and_tags,omitempty"`

	// 手动创建时的资源详情
	Resources *[]Resource `json:"resources,omitempty"`

	// 手动创建，选择资源层级为云产品时的资源详情
	ProductResources *[]ProductResource `json:"product_resources,omitempty"`

	// 实例名称匹配参数
	Instances *[]Instance `json:"instances,omitempty"`

	// 创建资源层级为云产品时的云产品的取值，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。多个云产品则用“;”隔开，如\"SERVICE.BMS,instance_id;SYS.ECS,instance_id\"。
	ProductNames *string `json:"product_names,omitempty"`

	// 资源层级，资源生效范围。选择云产品，则云产品及其子层级均可进入该资源分组，选择子维度，则只生效具体的子维度 product 云产品 dimension 子维度
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
