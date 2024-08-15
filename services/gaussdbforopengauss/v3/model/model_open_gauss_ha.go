package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussHa 实例部署形态。
type OpenGaussHa struct {

	// GaussDB为分布式时，取值：enterprise；为集中式时，取值：centralization_standard。不区分大小写。
	Mode OpenGaussHaMode `json:"mode"`

	// 指定实例一致性类型，当创建分布式模式实例时，该字段值必传，当创建主备模式实例时，该字段值不传。取值范围：strong（强一致性） | eventual(最终一致性)，不分区大小写。
	Consistency OpenGaussHaConsistency `json:"consistency"`

	// 备机同步参数。  取值：  GaussDB为“sync\"  说明： - “sync”为同步模式。
	ReplicationMode OpenGaussHaReplicationMode `json:"replication_mode"`

	// 指定创建实例的产品类型，创建企业版实例时传空值或者enterprise，创建基础版实例时需要指定instance_mode的值为basic，创建生态版实例时需要指定instance_mode的值为ecology。
	InstanceMode *OpenGaussHaInstanceMode `json:"instance_mode,omitempty"`
}

func (o OpenGaussHa) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussHa struct{}"
	}

	return strings.Join([]string{"OpenGaussHa", string(data)}, " ")
}

type OpenGaussHaMode struct {
	value string
}

type OpenGaussHaModeEnum struct {
	ENTERPRISE              OpenGaussHaMode
	CENTRALIZATION_STANDARD OpenGaussHaMode
}

func GetOpenGaussHaModeEnum() OpenGaussHaModeEnum {
	return OpenGaussHaModeEnum{
		ENTERPRISE: OpenGaussHaMode{
			value: "enterprise",
		},
		CENTRALIZATION_STANDARD: OpenGaussHaMode{
			value: "centralization_standard",
		},
	}
}

func (c OpenGaussHaMode) Value() string {
	return c.value
}

func (c OpenGaussHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaMode) UnmarshalJSON(b []byte) error {
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

func (c OpenGaussHaConsistency) Value() string {
	return c.value
}

func (c OpenGaussHaConsistency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaConsistency) UnmarshalJSON(b []byte) error {
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

func (c OpenGaussHaReplicationMode) Value() string {
	return c.value
}

func (c OpenGaussHaReplicationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaReplicationMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaInstanceMode struct {
	value string
}

type OpenGaussHaInstanceModeEnum struct {
	ENTERPRISE OpenGaussHaInstanceMode
	BASIC      OpenGaussHaInstanceMode
	ECOLOGY    OpenGaussHaInstanceMode
}

func GetOpenGaussHaInstanceModeEnum() OpenGaussHaInstanceModeEnum {
	return OpenGaussHaInstanceModeEnum{
		ENTERPRISE: OpenGaussHaInstanceMode{
			value: "enterprise",
		},
		BASIC: OpenGaussHaInstanceMode{
			value: "basic",
		},
		ECOLOGY: OpenGaussHaInstanceMode{
			value: "ecology",
		},
	}
}

func (c OpenGaussHaInstanceMode) Value() string {
	return c.value
}

func (c OpenGaussHaInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaInstanceMode) UnmarshalJSON(b []byte) error {
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
