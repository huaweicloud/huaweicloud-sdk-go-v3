package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BucketType struct {
	value string
}

type BucketTypeEnum struct {
	OBJECT BucketType
	PFS    BucketType
}

func GetBucketTypeEnum() BucketTypeEnum {
	return BucketTypeEnum{
		OBJECT: BucketType{
			value: "OBJECT",
		},
		PFS: BucketType{
			value: "PFS",
		},
	}
}

func (c BucketType) Value() string {
	return c.value
}

func (c BucketType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BucketType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
