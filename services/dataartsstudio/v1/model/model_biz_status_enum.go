package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BizStatusEnum 实体的发布状态，只读，创建和更新时无需填写。 枚举值：   - DRAFT: 草稿   - PUBLISH_DEVELOPING: 发布待审核   - PUBLISHED: 已发布   - OFFLINE_DEVELOPING: 下线待审核   - OFFLINE: 已下线   - REJECT: 已驳回
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
