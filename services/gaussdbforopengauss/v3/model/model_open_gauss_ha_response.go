package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussHaResponse 实例部署形态。
type OpenGaussHaResponse struct {

	// GaussDB 分布式模式，返回值为：Enterprise（企业版）；主备版，返回值为：Ha(主备版)。
	Mode OpenGaussHaResponseMode `json:"mode"`

	// 备机同步参数。  取值：  GaussDB为“sync”。 说明： - “sync”为同步模式。
	ReplicationMode OpenGaussHaResponseReplicationMode `json:"replication_mode"`

	// GaussDB的预留参数：指定实例一致性类型，取值范围：strong（强一致性） | eventual(最终一致性)。
	Consistency OpenGaussHaResponseConsistency `json:"consistency"`
}

func (o OpenGaussHaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussHaResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussHaResponse", string(data)}, " ")
}

type OpenGaussHaResponseMode struct {
	value string
}

type OpenGaussHaResponseModeEnum struct {
	ENTERPRISE OpenGaussHaResponseMode
	HA         OpenGaussHaResponseMode
}

func GetOpenGaussHaResponseModeEnum() OpenGaussHaResponseModeEnum {
	return OpenGaussHaResponseModeEnum{
		ENTERPRISE: OpenGaussHaResponseMode{
			value: "Enterprise",
		},
		HA: OpenGaussHaResponseMode{
			value: "Ha",
		},
	}
}

func (c OpenGaussHaResponseMode) Value() string {
	return c.value
}

func (c OpenGaussHaResponseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResponseMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaResponseReplicationMode struct {
	value string
}

type OpenGaussHaResponseReplicationModeEnum struct {
	SYNC OpenGaussHaResponseReplicationMode
}

func GetOpenGaussHaResponseReplicationModeEnum() OpenGaussHaResponseReplicationModeEnum {
	return OpenGaussHaResponseReplicationModeEnum{
		SYNC: OpenGaussHaResponseReplicationMode{
			value: "sync",
		},
	}
}

func (c OpenGaussHaResponseReplicationMode) Value() string {
	return c.value
}

func (c OpenGaussHaResponseReplicationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResponseReplicationMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaResponseConsistency struct {
	value string
}

type OpenGaussHaResponseConsistencyEnum struct {
	STRONG   OpenGaussHaResponseConsistency
	EVENTUAL OpenGaussHaResponseConsistency
}

func GetOpenGaussHaResponseConsistencyEnum() OpenGaussHaResponseConsistencyEnum {
	return OpenGaussHaResponseConsistencyEnum{
		STRONG: OpenGaussHaResponseConsistency{
			value: "strong",
		},
		EVENTUAL: OpenGaussHaResponseConsistency{
			value: "eventual",
		},
	}
}

func (c OpenGaussHaResponseConsistency) Value() string {
	return c.value
}

func (c OpenGaussHaResponseConsistency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResponseConsistency) UnmarshalJSON(b []byte) error {
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
