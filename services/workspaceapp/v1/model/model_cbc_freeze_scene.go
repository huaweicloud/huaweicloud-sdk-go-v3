package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CbcFreezeScene 更新云服务状态的业务场景，默认`ARREAR`。 * `ARREAR` - 欠费场景；为正常的运营业务场景，包括包周期资源到期、按需资源扣费失败 * `POLICE` - 公安冻结场景 * `ILLEGAL` - 违规冻结场景 * `VERIFY` - 客户未实名认证冻结场景 * `PARTNER` - 合作伙伴冻结（合作伙伴冻结子客户资源）
type CbcFreezeScene struct {
	value string
}

type CbcFreezeSceneEnum struct {
	ARREAR  CbcFreezeScene
	POLICE  CbcFreezeScene
	ILLEGAL CbcFreezeScene
	VERIFY  CbcFreezeScene
	PARTNER CbcFreezeScene
}

func GetCbcFreezeSceneEnum() CbcFreezeSceneEnum {
	return CbcFreezeSceneEnum{
		ARREAR: CbcFreezeScene{
			value: "ARREAR",
		},
		POLICE: CbcFreezeScene{
			value: "POLICE",
		},
		ILLEGAL: CbcFreezeScene{
			value: "ILLEGAL",
		},
		VERIFY: CbcFreezeScene{
			value: "VERIFY",
		},
		PARTNER: CbcFreezeScene{
			value: "PARTNER",
		},
	}
}

func (c CbcFreezeScene) Value() string {
	return c.value
}

func (c CbcFreezeScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CbcFreezeScene) UnmarshalJSON(b []byte) error {
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
