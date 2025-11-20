package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectionInfo 收藏项
type CollectionInfo struct {

	// 收藏项类型。 * VOICE：音色
	CollectionType *CollectionInfoCollectionType `json:"collection_type,omitempty"`

	// 收藏的资产或者模板ID。
	CollectionIds *[]string `json:"collection_ids,omitempty"`
}

func (o CollectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectionInfo struct{}"
	}

	return strings.Join([]string{"CollectionInfo", string(data)}, " ")
}

type CollectionInfoCollectionType struct {
	value string
}

type CollectionInfoCollectionTypeEnum struct {
	VOICE CollectionInfoCollectionType
}

func GetCollectionInfoCollectionTypeEnum() CollectionInfoCollectionTypeEnum {
	return CollectionInfoCollectionTypeEnum{
		VOICE: CollectionInfoCollectionType{
			value: "VOICE",
		},
	}
}

func (c CollectionInfoCollectionType) Value() string {
	return c.value
}

func (c CollectionInfoCollectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectionInfoCollectionType) UnmarshalJSON(b []byte) error {
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
