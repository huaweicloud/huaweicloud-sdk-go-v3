package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DomainIdFilterTypePrimitiveTypeHolder struct {

	// 租户ID筛选类型，仅支持资源栈集权限模型是SERVICE_MANAGED时指定该参数。默认为NONE  用户可以指定不同的筛选类型，通过domain_ids或domain_ids_uri 指定或排除部署的用户信息，以增加或限制部署目标用户范围，实现不同的部署策略。     * `INTERSECTION` - 从所提供的 OUs 中选择指定用户进行部署。指定该类型时，domain_ids和domain_ids_uri 有且仅能有一个存在。    * `DIFFERENCE` - 从所提供的 OUs 中排除指定用户部署。指定该类型时，domain_ids和domain_ids_uri 有且仅能有一个存在。    * `UNION` - 除了部署到提供的 OUs 外，还将部署到指定的账户中。用户可以通过同时指定organizational_unit_ids 和 domain_ids/domain_ids_uri 以实现在同一请求中部署整个OU和来自不同OU的特定个人账户。指定该类型时，domain_ids和domain_ids_uri 有且仅能有一个存在。创建资源栈实例（CreateStackInstances）除外，该API 不允许指定使用该类型。    * `NONE` - 只部署在指定 OUs 中的所有账户。该类型下domain_ids和domain_ids_uri字段必须为空。
	DomainIdFilterType *DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType `json:"domain_id_filter_type,omitempty"`
}

func (o DomainIdFilterTypePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainIdFilterTypePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"DomainIdFilterTypePrimitiveTypeHolder", string(data)}, " ")
}

type DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType struct {
	value string
}

type DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterTypeEnum struct {
	INTERSECTION DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType
	DIFFERENCE   DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType
	UNION        DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType
	NONE         DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType
}

func GetDomainIdFilterTypePrimitiveTypeHolderDomainIdFilterTypeEnum() DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterTypeEnum {
	return DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterTypeEnum{
		INTERSECTION: DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType{
			value: "INTERSECTION",
		},
		DIFFERENCE: DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType{
			value: "DIFFERENCE",
		},
		UNION: DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType{
			value: "UNION",
		},
		NONE: DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType{
			value: "NONE",
		},
	}
}

func (c DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType) Value() string {
	return c.value
}

func (c DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainIdFilterTypePrimitiveTypeHolderDomainIdFilterType) UnmarshalJSON(b []byte) error {
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
