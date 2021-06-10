package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 迁移实例信息体
type InstInfo struct {
	// 引擎类型

	EngineType *InstInfoEngineType `json:"engine_type,omitempty"`
	// 实例类型

	InstType *InstInfoInstType `json:"inst_type,omitempty"`
	// 迁移实例所在的私有IP

	Ip *string `json:"ip,omitempty"`
	// 迁移实例所在的公网IP

	PublicIp *string `json:"public_ip,omitempty"`
	// 迁移实例任务定时启动时间

	StartTime *string `json:"start_time,omitempty"`
	// 迁移实例的状态

	Status *InstInfoStatus `json:"status,omitempty"`
	// 迁移实例的磁盘大小

	VolumeSize *int32 `json:"volume_size,omitempty"`
}

func (o InstInfo) String() string {
	data, err := json.Marshal(o)
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
	}
}

func (c InstInfoEngineType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstInfoEngineType) UnmarshalJSON(b []byte) error {
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

func (c InstInfoInstType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstInfoInstType) UnmarshalJSON(b []byte) error {
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

func (c InstInfoStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstInfoStatus) UnmarshalJSON(b []byte) error {
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
