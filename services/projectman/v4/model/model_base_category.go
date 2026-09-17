package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseCategory Category基础信息
type BaseCategory struct {

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释**： 修改人。 **取值范围**： 不涉及。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// **参数解释**： 修改时间。 **取值范围**： 不涉及。
	ModifiedDate *string `json:"modified_date,omitempty"`

	// **参数解释**： 创建人。 **取值范围**： 不涉及。
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreatedDate *string `json:"created_date,omitempty"`

	// **参数解释**： 对象类型编码。 **取值范围**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 编号前缀。 **取值范围**： 不涉及。
	Prefix *string `json:"prefix,omitempty"`

	// **参数解释**： 租户下项目空间唯一标识ID。 **取值范围**： - -1：自定义工作项类型 - 0：预设工作项模型
	DomainId *BaseCategoryDomainId `json:"domain_id,omitempty"`

	// **参数解释**： 图标。 **取值范围**： 不涉及。
	Icon *string `json:"icon,omitempty"`

	// **参数解释**： 颜色。 **取值范围**： 不涉及。
	Color *string `json:"color,omitempty"`

	// **参数解释**： 描述信息。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 定义类型。 **取值范围**： - 1~3 系统级别 - 4 租户级别
	DefinitionType *int64 `json:"definition_type,omitempty"`

	// **参数解释**： 类别ID。 **取值范围**： 不涉及。
	TypeId *string `json:"type_id,omitempty"`
}

func (o BaseCategory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseCategory struct{}"
	}

	return strings.Join([]string{"BaseCategory", string(data)}, " ")
}

type BaseCategoryDomainId struct {
	value string
}

type BaseCategoryDomainIdEnum struct {
	E_1 BaseCategoryDomainId
	E_0 BaseCategoryDomainId
}

func GetBaseCategoryDomainIdEnum() BaseCategoryDomainIdEnum {
	return BaseCategoryDomainIdEnum{
		E_1: BaseCategoryDomainId{
			value: "-1",
		},
		E_0: BaseCategoryDomainId{
			value: "0",
		},
	}
}

func (c BaseCategoryDomainId) Value() string {
	return c.value
}

func (c BaseCategoryDomainId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseCategoryDomainId) UnmarshalJSON(b []byte) error {
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
