package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServerStatus 实例的状态 * `UNREGISTER` - 未就绪 * `REGISTERED` - 就绪状态 * `MAINTAINING` - 维护中 * `FREEZE` - 冻结 * `STOPPED` - 停止 * `NONE` - 异常状态
type ServerStatus struct {
	value string
}

type ServerStatusEnum struct {
	UNREGISTER  ServerStatus
	REGISTERED  ServerStatus
	MAINTAINING ServerStatus
	FREEZE      ServerStatus
	STOPPED     ServerStatus
	NONE        ServerStatus
}

func GetServerStatusEnum() ServerStatusEnum {
	return ServerStatusEnum{
		UNREGISTER: ServerStatus{
			value: "UNREGISTER",
		},
		REGISTERED: ServerStatus{
			value: "REGISTERED",
		},
		MAINTAINING: ServerStatus{
			value: "MAINTAINING",
		},
		FREEZE: ServerStatus{
			value: "FREEZE",
		},
		STOPPED: ServerStatus{
			value: "STOPPED",
		},
		NONE: ServerStatus{
			value: "NONE",
		},
	}
}

func (c ServerStatus) Value() string {
	return c.value
}

func (c ServerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerStatus) UnmarshalJSON(b []byte) error {
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
