package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NamespaceMetadata struct {

	// 是否公开，可选true、false
	Public NamespaceMetadataPublic `json:"public"`
}

func (o NamespaceMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceMetadata struct{}"
	}

	return strings.Join([]string{"NamespaceMetadata", string(data)}, " ")
}

type NamespaceMetadataPublic struct {
	value string
}

type NamespaceMetadataPublicEnum struct {
	TRUE  NamespaceMetadataPublic
	FALSE NamespaceMetadataPublic
}

func GetNamespaceMetadataPublicEnum() NamespaceMetadataPublicEnum {
	return NamespaceMetadataPublicEnum{
		TRUE: NamespaceMetadataPublic{
			value: "true",
		},
		FALSE: NamespaceMetadataPublic{
			value: "false",
		},
	}
}

func (c NamespaceMetadataPublic) Value() string {
	return c.value
}

func (c NamespaceMetadataPublic) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NamespaceMetadataPublic) UnmarshalJSON(b []byte) error {
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
