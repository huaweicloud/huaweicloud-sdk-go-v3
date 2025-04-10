package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Scene 资源运营状态场景： ARREAR：欠费场景； POLICE：公安场景。 ILLEGAL：违规场景。 VERIFY：客户未实名认证冻结场景。
type Scene struct {
	value string
}

type SceneEnum struct {
	ARREAR  Scene
	POLICE  Scene
	ILLEGAL Scene
	VERIFY  Scene
}

func GetSceneEnum() SceneEnum {
	return SceneEnum{
		ARREAR: Scene{
			value: "ARREAR",
		},
		POLICE: Scene{
			value: "POLICE",
		},
		ILLEGAL: Scene{
			value: "ILLEGAL",
		},
		VERIFY: Scene{
			value: "VERIFY",
		},
	}
}

func (c Scene) Value() string {
	return c.value
}

func (c Scene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Scene) UnmarshalJSON(b []byte) error {
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
