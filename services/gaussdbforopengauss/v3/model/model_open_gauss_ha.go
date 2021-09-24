package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建分布式实例时使用。
type OpenGaussHa struct {
	// GaussDB(for openGauss)为分布式模式，取值：enterprise(企业版) ，不区分大小写。

	Mode OpenGaussHaMode `json:"mode"`
	// 备机同步参数。  取值：  GaussDB(for openGauss)为“sync\"  说明： - “sync”为同步模式。

	ReplicationMode OpenGaussHaReplicationMode `json:"replication_mode"`
	// 指定实例一致性类型，取值范围：strong（强一致性） | eventual(最终一致性)，不分区大小写。

	Consistency OpenGaussHaConsistency `json:"consistency"`
}

func (o OpenGaussHa) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpenGaussHa struct{}"
	}

	return strings.Join([]string{"OpenGaussHa", string(data)}, " ")
}

type OpenGaussHaMode struct {
	value string
}

type OpenGaussHaModeEnum struct {
	ENTERPRISE OpenGaussHaMode
}

func GetOpenGaussHaModeEnum() OpenGaussHaModeEnum {
	return OpenGaussHaModeEnum{
		ENTERPRISE: OpenGaussHaMode{
			value: "enterprise",
		},
	}
}

func (c OpenGaussHaMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *OpenGaussHaMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaReplicationMode struct {
	value string
}

type OpenGaussHaReplicationModeEnum struct {
	SYNC OpenGaussHaReplicationMode
}

func GetOpenGaussHaReplicationModeEnum() OpenGaussHaReplicationModeEnum {
	return OpenGaussHaReplicationModeEnum{
		SYNC: OpenGaussHaReplicationMode{
			value: "sync",
		},
	}
}

func (c OpenGaussHaReplicationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *OpenGaussHaReplicationMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaConsistency struct {
	value string
}

type OpenGaussHaConsistencyEnum struct {
	STRONG   OpenGaussHaConsistency
	EVENTUAL OpenGaussHaConsistency
}

func GetOpenGaussHaConsistencyEnum() OpenGaussHaConsistencyEnum {
	return OpenGaussHaConsistencyEnum{
		STRONG: OpenGaussHaConsistency{
			value: "strong",
		},
		EVENTUAL: OpenGaussHaConsistency{
			value: "eventual",
		},
	}
}

func (c OpenGaussHaConsistency) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *OpenGaussHaConsistency) UnmarshalJSON(b []byte) error {
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
