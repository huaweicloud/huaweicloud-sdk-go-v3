package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChardMode 资源计费类型。 * ON_DEMAND:按需计费，目前只有进阶版声音，最多制作三个任务 * PERIODIC: 包周期 * ONE_TIME：一次性计费 > * 一次性计费包括：租户订购的一次性资源，SP管理员分配给租户的一次性资源。
type ChardMode struct {
	value string
}

type ChardModeEnum struct {
	PERIODIC  ChardMode
	ONE_TIME  ChardMode
	ON_DEMAND ChardMode
}

func GetChardModeEnum() ChardModeEnum {
	return ChardModeEnum{
		PERIODIC: ChardMode{
			value: "PERIODIC",
		},
		ONE_TIME: ChardMode{
			value: "ONE_TIME",
		},
		ON_DEMAND: ChardMode{
			value: "ON_DEMAND",
		},
	}
}

func (c ChardMode) Value() string {
	return c.value
}

func (c ChardMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChardMode) UnmarshalJSON(b []byte) error {
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
