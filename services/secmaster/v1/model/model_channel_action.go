package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChannelAction **参数解释**: 节点运行状态的监控 - START 开始 - STOP 停止 - REMOVE 移除 - RESTART 重启 - REFRESH 刷新 - INSTALL 安装  **约束限制** 不涉及 **取值范围**: - START - STOP - REMOVE - RESTART - REFRESH - INSTALL  **默认值** 不涉及
type ChannelAction struct {
	value string
}

type ChannelActionEnum struct {
	START   ChannelAction
	STOP    ChannelAction
	REMOVE  ChannelAction
	RESTART ChannelAction
	REFRESH ChannelAction
	INSTALL ChannelAction
}

func GetChannelActionEnum() ChannelActionEnum {
	return ChannelActionEnum{
		START: ChannelAction{
			value: "START",
		},
		STOP: ChannelAction{
			value: "STOP",
		},
		REMOVE: ChannelAction{
			value: "REMOVE",
		},
		RESTART: ChannelAction{
			value: "RESTART",
		},
		REFRESH: ChannelAction{
			value: "REFRESH",
		},
		INSTALL: ChannelAction{
			value: "INSTALL",
		},
	}
}

func (c ChannelAction) Value() string {
	return c.value
}

func (c ChannelAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChannelAction) UnmarshalJSON(b []byte) error {
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
