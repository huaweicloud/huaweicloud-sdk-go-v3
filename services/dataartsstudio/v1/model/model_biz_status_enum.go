package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BizStatusEnum 业务状态
type BizStatusEnum struct {
	value string
}

type BizStatusEnumEnum struct {
	DRAFT              BizStatusEnum
	PUBLISH_DEVELOPING BizStatusEnum
	PUBLISHED          BizStatusEnum
	OFFLINE_DEVELOPING BizStatusEnum
	OFFLINE            BizStatusEnum
	REJECT             BizStatusEnum
}

func GetBizStatusEnumEnum() BizStatusEnumEnum {
	return BizStatusEnumEnum{
		DRAFT: BizStatusEnum{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: BizStatusEnum{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: BizStatusEnum{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: BizStatusEnum{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: BizStatusEnum{
			value: "OFFLINE",
		},
		REJECT: BizStatusEnum{
			value: "REJECT",
		},
	}
}

func (c BizStatusEnum) Value() string {
	return c.value
}

func (c BizStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BizStatusEnum) UnmarshalJSON(b []byte) error {
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
