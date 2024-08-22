package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussHaOption 实例部署形态。
type OpenGaussHaOption struct {

	// GaussDB为分布式时，取值：enterprise；为集中式时，取值：centralization_standard。不区分大小写。
	Mode OpenGaussHaOptionMode `json:"mode"`

	// 指定实例一致性类型，当创建分布式模式实例时，该字段值必传，当创建主备模式实例时，该字段值不传。取值范围：strong（强一致性） | eventual(最终一致性)，不分区大小写。
	Consistency OpenGaussHaOptionConsistency `json:"consistency"`

	// 备机同步参数。  取值：  GaussDB为“sync\"  说明： - “sync”为同步模式。
	ReplicationMode OpenGaussHaOptionReplicationMode `json:"replication_mode"`

	// 指定创建实例的产品类型，创建企业版实例时传空值或者enterprise，创建基础版实例时需要指定instance_mode的值为basic，创建生态版实例时需要指定instance_mode的值为ecology。
	InstanceMode *OpenGaussHaOptionInstanceMode `json:"instance_mode,omitempty"`
}

func (o OpenGaussHaOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussHaOption struct{}"
	}

	return strings.Join([]string{"OpenGaussHaOption", string(data)}, " ")
}

type OpenGaussHaOptionMode struct {
	value string
}

type OpenGaussHaOptionModeEnum struct {
	ENTERPRISE              OpenGaussHaOptionMode
	CENTRALIZATION_STANDARD OpenGaussHaOptionMode
}

func GetOpenGaussHaOptionModeEnum() OpenGaussHaOptionModeEnum {
	return OpenGaussHaOptionModeEnum{
		ENTERPRISE: OpenGaussHaOptionMode{
			value: "enterprise",
		},
		CENTRALIZATION_STANDARD: OpenGaussHaOptionMode{
			value: "centralization_standard",
		},
	}
}

func (c OpenGaussHaOptionMode) Value() string {
	return c.value
}

func (c OpenGaussHaOptionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaOptionMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaOptionConsistency struct {
	value string
}

type OpenGaussHaOptionConsistencyEnum struct {
	STRONG   OpenGaussHaOptionConsistency
	EVENTUAL OpenGaussHaOptionConsistency
}

func GetOpenGaussHaOptionConsistencyEnum() OpenGaussHaOptionConsistencyEnum {
	return OpenGaussHaOptionConsistencyEnum{
		STRONG: OpenGaussHaOptionConsistency{
			value: "strong",
		},
		EVENTUAL: OpenGaussHaOptionConsistency{
			value: "eventual",
		},
	}
}

func (c OpenGaussHaOptionConsistency) Value() string {
	return c.value
}

func (c OpenGaussHaOptionConsistency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaOptionConsistency) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaOptionReplicationMode struct {
	value string
}

type OpenGaussHaOptionReplicationModeEnum struct {
	SYNC OpenGaussHaOptionReplicationMode
}

func GetOpenGaussHaOptionReplicationModeEnum() OpenGaussHaOptionReplicationModeEnum {
	return OpenGaussHaOptionReplicationModeEnum{
		SYNC: OpenGaussHaOptionReplicationMode{
			value: "sync",
		},
	}
}

func (c OpenGaussHaOptionReplicationMode) Value() string {
	return c.value
}

func (c OpenGaussHaOptionReplicationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaOptionReplicationMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussHaOptionInstanceMode struct {
	value string
}

type OpenGaussHaOptionInstanceModeEnum struct {
	ENTERPRISE OpenGaussHaOptionInstanceMode
	BASIC      OpenGaussHaOptionInstanceMode
	ECOLOGY    OpenGaussHaOptionInstanceMode
}

func GetOpenGaussHaOptionInstanceModeEnum() OpenGaussHaOptionInstanceModeEnum {
	return OpenGaussHaOptionInstanceModeEnum{
		ENTERPRISE: OpenGaussHaOptionInstanceMode{
			value: "enterprise",
		},
		BASIC: OpenGaussHaOptionInstanceMode{
			value: "basic",
		},
		ECOLOGY: OpenGaussHaOptionInstanceMode{
			value: "ecology",
		},
	}
}

func (c OpenGaussHaOptionInstanceMode) Value() string {
	return c.value
}

func (c OpenGaussHaOptionInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussHaOptionInstanceMode) UnmarshalJSON(b []byte) error {
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
