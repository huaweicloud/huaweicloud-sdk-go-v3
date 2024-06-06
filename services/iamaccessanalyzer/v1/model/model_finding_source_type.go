package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FindingSourceType 访问分析结果的来源。
type FindingSourceType struct {
	value string
}

type FindingSourceTypeEnum struct {
	BUCKET_ACL    FindingSourceType
	BUCKET_POLICY FindingSourceType
}

func GetFindingSourceTypeEnum() FindingSourceTypeEnum {
	return FindingSourceTypeEnum{
		BUCKET_ACL: FindingSourceType{
			value: "bucket_acl",
		},
		BUCKET_POLICY: FindingSourceType{
			value: "bucket_policy",
		},
	}
}

func (c FindingSourceType) Value() string {
	return c.value
}

func (c FindingSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FindingSourceType) UnmarshalJSON(b []byte) error {
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
