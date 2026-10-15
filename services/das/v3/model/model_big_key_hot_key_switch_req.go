package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BigKeyHotKeySwitchReq struct {

	// 操作类型。enable-开启，disable-关闭
	Action BigKeyHotKeySwitchReqAction `json:"action"`

	// 引擎类型（小写），如geminidb-redis
	EngineType string `json:"engine_type"`
}

func (o BigKeyHotKeySwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BigKeyHotKeySwitchReq struct{}"
	}

	return strings.Join([]string{"BigKeyHotKeySwitchReq", string(data)}, " ")
}

type BigKeyHotKeySwitchReqAction struct {
	value string
}

type BigKeyHotKeySwitchReqActionEnum struct {
	ENABLE  BigKeyHotKeySwitchReqAction
	DISABLE BigKeyHotKeySwitchReqAction
}

func GetBigKeyHotKeySwitchReqActionEnum() BigKeyHotKeySwitchReqActionEnum {
	return BigKeyHotKeySwitchReqActionEnum{
		ENABLE: BigKeyHotKeySwitchReqAction{
			value: "enable",
		},
		DISABLE: BigKeyHotKeySwitchReqAction{
			value: "disable",
		},
	}
}

func (c BigKeyHotKeySwitchReqAction) Value() string {
	return c.value
}

func (c BigKeyHotKeySwitchReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BigKeyHotKeySwitchReqAction) UnmarshalJSON(b []byte) error {
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
