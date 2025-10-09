package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PackageOptions 配置参数结构
type PackageOptions struct {

	// **参数解释**： 参数名称 **取值范围**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 参数默认值，不指定时按默认值生效, 参数类型以实际返回为准，可能为integer,string或者boolean **取值范围**： 不涉及
	Default *interface{} `json:"default"`

	// **参数解释**： 参数生效方式 **取值范围**： - static：集群、节点池创建时生效，后续不可修改 - immediately：集群、节点池运行中时可以修改，修改后生效
	ValidAt PackageOptionsValidAt `json:"validAt"`

	// **参数解释**： 配置项是否可以为空 **取值范围**： - true：配置项为空时，不使用默认值，为空值 - false：配置项为空时，使用默认值
	Empty bool `json:"empty"`

	// **参数解释**： 参数分类 **取值范围**： 不涉及
	Schema string `json:"schema"`

	// **参数解释**： 参数类型 **取值范围**： 不涉及
	Type string `json:"type"`
}

func (o PackageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageOptions struct{}"
	}

	return strings.Join([]string{"PackageOptions", string(data)}, " ")
}

type PackageOptionsValidAt struct {
	value string
}

type PackageOptionsValidAtEnum struct {
	STATIC      PackageOptionsValidAt
	IMMEDIATELY PackageOptionsValidAt
}

func GetPackageOptionsValidAtEnum() PackageOptionsValidAtEnum {
	return PackageOptionsValidAtEnum{
		STATIC: PackageOptionsValidAt{
			value: "static",
		},
		IMMEDIATELY: PackageOptionsValidAt{
			value: "immediately",
		},
	}
}

func (c PackageOptionsValidAt) Value() string {
	return c.value
}

func (c PackageOptionsValidAt) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PackageOptionsValidAt) UnmarshalJSON(b []byte) error {
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
