package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 实例部署形态。
type OpenGaussHaResult struct {

	// GaussDB 分布式模式，返回值为：enterprise（企业版）；主备版，返回值为：centralization_standard(主备版)。
	Mode OpenGaussHaResultMode `json:"mode"`

	// 备机同步参数。  取值：  GaussDB为“sync”。 说明： - “sync”为同步模式。
	ReplicationMode OpenGaussHaResultReplicationMode `json:"replication_mode"`

	// GaussDB的预留参数：指定实例一致性类型，取值范围：strong（强一致性） | eventual(最终一致性)。
	Consistency OpenGaussHaResultConsistency `json:"consistency"`

	// 指定副本一致性协议类型，取值范围：quorum | paxos。不填时，默认为quorum。
	ConsistencyProtocol *string `json:"consistency_protocol,omitempty"`
}

func (o OpenGaussHaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussHaResult struct{}"
	}

	return strings.Join([]string{"OpenGaussHaResult", string(data)}, " ")
}

type OpenGaussHaResultMode struct {
	value string
}

type OpenGaussHaResultModeEnum struct {
	ENTERPRISE              OpenGaussHaResultMode
	CENTRALIZATION_STANDARD OpenGaussHaResultMode
}

func GetOpenGaussHaResultModeEnum() OpenGaussHaResultModeEnum {
	return OpenGaussHaResultModeEnum{
		ENTERPRISE: OpenGaussHaResultMode{
			value: "Enterprise",
		},
		CENTRALIZATION_STANDARD: OpenGaussHaResultMode{
			value: "centralization_standard",
		},
	}
}

func (c OpenGaussHaResultMode) Value() string {
	return c.value
}

func (c OpenGaussHaResultMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResultMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaResultReplicationMode struct {
	value string
}

type OpenGaussHaResultReplicationModeEnum struct {
	SYNC OpenGaussHaResultReplicationMode
}

func GetOpenGaussHaResultReplicationModeEnum() OpenGaussHaResultReplicationModeEnum {
	return OpenGaussHaResultReplicationModeEnum{
		SYNC: OpenGaussHaResultReplicationMode{
			value: "sync",
		},
	}
}

func (c OpenGaussHaResultReplicationMode) Value() string {
	return c.value
}

func (c OpenGaussHaResultReplicationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResultReplicationMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaResultConsistency struct {
	value string
}

type OpenGaussHaResultConsistencyEnum struct {
	STRONG   OpenGaussHaResultConsistency
	EVENTUAL OpenGaussHaResultConsistency
}

func GetOpenGaussHaResultConsistencyEnum() OpenGaussHaResultConsistencyEnum {
	return OpenGaussHaResultConsistencyEnum{
		STRONG: OpenGaussHaResultConsistency{
			value: "strong",
		},
		EVENTUAL: OpenGaussHaResultConsistency{
			value: "eventual",
		},
	}
}

func (c OpenGaussHaResultConsistency) Value() string {
	return c.value
}

func (c OpenGaussHaResultConsistency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaResultConsistency) UnmarshalJSON(b []byte) error {
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
