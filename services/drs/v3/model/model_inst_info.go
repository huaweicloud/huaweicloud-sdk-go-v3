package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstInfo 迁移实例信息体
type InstInfo struct {

	// **参数解释**：  DRS任务引擎类型。  **约束限制**：  不涉及。  **取值范围**：  - mysql：MySQL到MySQL迁移，MySQL到MySQL同步 - mongodb：MongoDB到DDS迁移与同步 - cloudDataGuard-mysql：MySQL到MySQL灾备 - gaussdbv5：GaussDB同步 - postgresql：PostgreSQL到PostgreSQL同步 - mysql-to-kafka：MySQL到Kafka同步 - taurus-to-kafka：TaurusDB到Kafka同步 - gaussdbv5ha-to-kafka：GaussDB集中式版到Kafka同步  **默认取值**：  不涉及
	EngineType *InstInfoEngineType `json:"engine_type,omitempty"`

	// **参数解释**：  实例规格类型。  **约束限制**：  不涉及。  **取值范围**：  - common：中 - high：大 - ultrahigh：极大  **默认取值**：  不涉及
	InstType *InstInfoInstType `json:"inst_type,omitempty"`

	// **参数解释**：  迁移实例所在的私有IP地址。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Ip *string `json:"ip,omitempty"`

	// **参数解释**：  迁移实例所在的公网IP地址。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**：  迁移实例任务定时启动时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  迁移实例的状态。  **约束限制**：  不涉及。  **取值范围**：  - active：运行中 - deleted：已删除  **默认取值**：  不涉及
	Status *InstInfoStatus `json:"status,omitempty"`

	// **参数解释**：  迁移实例的磁盘大小，单位为GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	VolumeSize *int32 `json:"volume_size,omitempty"`
}

func (o InstInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstInfo struct{}"
	}

	return strings.Join([]string{"InstInfo", string(data)}, " ")
}

type InstInfoEngineType struct {
	value string
}

type InstInfoEngineTypeEnum struct {
	MYSQL                  InstInfoEngineType
	MONGODB                InstInfoEngineType
	CLOUD_DATA_GUARD_MYSQL InstInfoEngineType
	GAUSSDBV5              InstInfoEngineType
	POSTGRESQL             InstInfoEngineType
	MYSQL_TO_KAFKA         InstInfoEngineType
	TAURUS_TO_KAFKA        InstInfoEngineType
	GAUSSDBV5HA_TO_KAFKA   InstInfoEngineType
}

func GetInstInfoEngineTypeEnum() InstInfoEngineTypeEnum {
	return InstInfoEngineTypeEnum{
		MYSQL: InstInfoEngineType{
			value: "mysql",
		},
		MONGODB: InstInfoEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: InstInfoEngineType{
			value: "cloudDataGuard-mysql",
		},
		GAUSSDBV5: InstInfoEngineType{
			value: "gaussdbv5",
		},
		POSTGRESQL: InstInfoEngineType{
			value: "postgresql",
		},
		MYSQL_TO_KAFKA: InstInfoEngineType{
			value: "mysql-to-kafka",
		},
		TAURUS_TO_KAFKA: InstInfoEngineType{
			value: "taurus-to-kafka",
		},
		GAUSSDBV5HA_TO_KAFKA: InstInfoEngineType{
			value: "gaussdbv5ha-to-kafka",
		},
	}
}

func (c InstInfoEngineType) Value() string {
	return c.value
}

func (c InstInfoEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstInfoEngineType) UnmarshalJSON(b []byte) error {
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

type InstInfoInstType struct {
	value string
}

type InstInfoInstTypeEnum struct {
	HIGH InstInfoInstType
}

func GetInstInfoInstTypeEnum() InstInfoInstTypeEnum {
	return InstInfoInstTypeEnum{
		HIGH: InstInfoInstType{
			value: "high",
		},
	}
}

func (c InstInfoInstType) Value() string {
	return c.value
}

func (c InstInfoInstType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstInfoInstType) UnmarshalJSON(b []byte) error {
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

type InstInfoStatus struct {
	value string
}

type InstInfoStatusEnum struct {
	ACTIVE  InstInfoStatus
	DELETED InstInfoStatus
}

func GetInstInfoStatusEnum() InstInfoStatusEnum {
	return InstInfoStatusEnum{
		ACTIVE: InstInfoStatus{
			value: "active",
		},
		DELETED: InstInfoStatus{
			value: "deleted",
		},
	}
}

func (c InstInfoStatus) Value() string {
	return c.value
}

func (c InstInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstInfoStatus) UnmarshalJSON(b []byte) error {
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
