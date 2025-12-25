package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HeartBeat **参数解释**: 节点是否成功收到心跳信号 - ONLINE 在线 - OFFLINE 离线  **约束限制** 不涉及 **取值范围**: - ONLINE - OFFLINE  **默认值** 不涉及
type HeartBeat struct {
	value string
}

type HeartBeatEnum struct {
	ONLINE  HeartBeat
	OFFLINE HeartBeat
}

func GetHeartBeatEnum() HeartBeatEnum {
	return HeartBeatEnum{
		ONLINE: HeartBeat{
			value: "ONLINE",
		},
		OFFLINE: HeartBeat{
			value: "OFFLINE",
		},
	}
}

func (c HeartBeat) Value() string {
	return c.value
}

func (c HeartBeat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HeartBeat) UnmarshalJSON(b []byte) error {
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
