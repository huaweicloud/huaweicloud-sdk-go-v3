package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建分布式实例时可见。
type OpenGaussHaResponse struct {
	// GaussDB(for openGauss)仅支持分布式模式，返回值为：Enterprise（企业版）。

	Mode OpenGaussHaResponseMode `json:"mode"`
	// 备机同步参数。  取值：  GaussDB(for openGauss)为“sync”。 说明： - “sync”为同步模式。

	ReplicationMode OpenGaussHaResponseReplicationMode `json:"replication_mode"`
	// GaussDB(for openGauss)的预留参数：指定实例一致性类型，取值范围：strong（强一致性） | eventual(最终一致性)。

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
}

func GetOpenGaussHaResponseModeEnum() OpenGaussHaResponseModeEnum {
	return OpenGaussHaResponseModeEnum{
		ENTERPRISE: OpenGaussHaResponseMode{
			value: "Enterprise",
		},
	}
}

func (c OpenGaussHaResponseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResponseMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
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

func (c OpenGaussHaResponseReplicationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResponseReplicationMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
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

func (c OpenGaussHaResponseConsistency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResponseConsistency) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
