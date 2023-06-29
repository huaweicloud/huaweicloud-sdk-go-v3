package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceInfo catalog input when grant policy
type ResourceInfo struct {

	// catalog info
	Catalogs *[]CatalogInfo `json:"catalogs,omitempty"`

	// uri
	Uris *[]string `json:"uris,omitempty"`

	// resource type
	Type ResourceInfoType `json:"type"`
}

func (o ResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInfo struct{}"
	}

	return strings.Join([]string{"ResourceInfo", string(data)}, " ")
}

type ResourceInfoType struct {
	value string
}

type ResourceInfoTypeEnum struct {
	CATALOG  ResourceInfoType
	DATABASE ResourceInfoType
	TABLE    ResourceInfoType
	COLUMN   ResourceInfoType
	FUNC     ResourceInfoType
	MODEL    ResourceInfoType
	URI      ResourceInfoType
}

func GetResourceInfoTypeEnum() ResourceInfoTypeEnum {
	return ResourceInfoTypeEnum{
		CATALOG: ResourceInfoType{
			value: "CATALOG",
		},
		DATABASE: ResourceInfoType{
			value: "DATABASE",
		},
		TABLE: ResourceInfoType{
			value: "TABLE",
		},
		COLUMN: ResourceInfoType{
			value: "COLUMN",
		},
		FUNC: ResourceInfoType{
			value: "FUNC",
		},
		MODEL: ResourceInfoType{
			value: "MODEL",
		},
		URI: ResourceInfoType{
			value: "URI",
		},
	}
}

func (c ResourceInfoType) Value() string {
	return c.value
}

func (c ResourceInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceInfoType) UnmarshalJSON(b []byte) error {
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
