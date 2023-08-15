package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentSubCategory 应用组件子类型。  Webapp的子类型有Web、Magento、Wordpress。  MicroService的子类型有Java Chassis、Go Chassis、Mesher、SpringCloud。  Common的子类型可以为空。
type ComponentSubCategory struct {
	value string
}

type ComponentSubCategoryEnum struct {
	WEB          ComponentSubCategory
	MAGENTO      ComponentSubCategory
	WORDPRESS    ComponentSubCategory
	SPRING_CLOUD ComponentSubCategory
	JAVA_CHASSIS ComponentSubCategory
	GO_CHASSIS   ComponentSubCategory
	MESHER       ComponentSubCategory
}

func GetComponentSubCategoryEnum() ComponentSubCategoryEnum {
	return ComponentSubCategoryEnum{
		WEB: ComponentSubCategory{
			value: "Web",
		},
		MAGENTO: ComponentSubCategory{
			value: "Magento",
		},
		WORDPRESS: ComponentSubCategory{
			value: "Wordpress",
		},
		SPRING_CLOUD: ComponentSubCategory{
			value: "SpringCloud",
		},
		JAVA_CHASSIS: ComponentSubCategory{
			value: "Java Chassis",
		},
		GO_CHASSIS: ComponentSubCategory{
			value: "Go Chassis",
		},
		MESHER: ComponentSubCategory{
			value: "Mesher",
		},
	}
}

func (c ComponentSubCategory) Value() string {
	return c.value
}

func (c ComponentSubCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentSubCategory) UnmarshalJSON(b []byte) error {
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
