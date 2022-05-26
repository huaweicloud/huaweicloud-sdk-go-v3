package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 子任务信息体
type ChildrenJobInfo struct {

	// 计费字段
	BillingTag bool `json:"billing_tag"`

	// 任务创建时间
	CreateTime string `json:"create_time"`

	// 复制场景
	DbUseType ChildrenJobInfoDbUseType `json:"db_use_type"`

	// 任务描述
	Description string `json:"description"`

	// 引擎类型
	EngineType ChildrenJobInfoEngineType `json:"engine_type"`

	// 任务失败原因
	ErrorMsg string `json:"error_msg"`

	// 任务id
	Id string `json:"id"`

	// 迁移方向
	JobDirection ChildrenJobInfoJobDirection `json:"job_direction"`

	// 任务名称
	Name string `json:"name"`

	// 网络类型
	NetType ChildrenJobInfoNetType `json:"net_type"`

	// 新框架
	NodeNewFramework bool `json:"node_newFramework"`

	// 任务状态
	Status string `json:"status"`

	// 迁移模式
	TaskType ChildrenJobInfoTaskType `json:"task_type"`
}

func (o ChildrenJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChildrenJobInfo struct{}"
	}

	return strings.Join([]string{"ChildrenJobInfo", string(data)}, " ")
}

type ChildrenJobInfoDbUseType struct {
	value string
}

type ChildrenJobInfoDbUseTypeEnum struct {
	MIGRATION        ChildrenJobInfoDbUseType
	SYNC             ChildrenJobInfoDbUseType
	CLOUD_DATA_GUARD ChildrenJobInfoDbUseType
}

func GetChildrenJobInfoDbUseTypeEnum() ChildrenJobInfoDbUseTypeEnum {
	return ChildrenJobInfoDbUseTypeEnum{
		MIGRATION: ChildrenJobInfoDbUseType{
			value: "migration:实时迁移",
		},
		SYNC: ChildrenJobInfoDbUseType{
			value: "sync:实时同步",
		},
		CLOUD_DATA_GUARD: ChildrenJobInfoDbUseType{
			value: "cloudDataGuard:实时灾备",
		},
	}
}

func (c ChildrenJobInfoDbUseType) Value() string {
	return c.value
}

func (c ChildrenJobInfoDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoDbUseType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoEngineType struct {
	value string
}

type ChildrenJobInfoEngineTypeEnum struct {
	CLOUD_DATA_GUARD_CASSANDRA       ChildrenJobInfoEngineType
	CLOUD_DATA_GUARD_DDM             ChildrenJobInfoEngineType
	CLOUD_DATA_GUARD_TAURUS_TO_MYSQL ChildrenJobInfoEngineType
	CLOUD_DATA_GUARD_MYSQL           ChildrenJobInfoEngineType
	CLOUD_DATA_GUARD_MYSQL_TO_TAURUS ChildrenJobInfoEngineType
}

func GetChildrenJobInfoEngineTypeEnum() ChildrenJobInfoEngineTypeEnum {
	return ChildrenJobInfoEngineTypeEnum{
		CLOUD_DATA_GUARD_CASSANDRA: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-cassandra",
		},
		CLOUD_DATA_GUARD_DDM: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-ddm",
		},
		CLOUD_DATA_GUARD_TAURUS_TO_MYSQL: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-taurus-to-mysql",
		},
		CLOUD_DATA_GUARD_MYSQL: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-mysql",
		},
		CLOUD_DATA_GUARD_MYSQL_TO_TAURUS: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-mysql-to-taurus",
		},
	}
}

func (c ChildrenJobInfoEngineType) Value() string {
	return c.value
}

func (c ChildrenJobInfoEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoEngineType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoJobDirection struct {
	value string
}

type ChildrenJobInfoJobDirectionEnum struct {
	UP     ChildrenJobInfoJobDirection
	DOWN   ChildrenJobInfoJobDirection
	NO_DBS ChildrenJobInfoJobDirection
}

func GetChildrenJobInfoJobDirectionEnum() ChildrenJobInfoJobDirectionEnum {
	return ChildrenJobInfoJobDirectionEnum{
		UP: ChildrenJobInfoJobDirection{
			value: "up",
		},
		DOWN: ChildrenJobInfoJobDirection{
			value: "down",
		},
		NO_DBS: ChildrenJobInfoJobDirection{
			value: "no-dbs",
		},
	}
}

func (c ChildrenJobInfoJobDirection) Value() string {
	return c.value
}

func (c ChildrenJobInfoJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoJobDirection) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoNetType struct {
	value string
}

type ChildrenJobInfoNetTypeEnum struct {
	VPC ChildrenJobInfoNetType
	VPN ChildrenJobInfoNetType
	EIP ChildrenJobInfoNetType
}

func GetChildrenJobInfoNetTypeEnum() ChildrenJobInfoNetTypeEnum {
	return ChildrenJobInfoNetTypeEnum{
		VPC: ChildrenJobInfoNetType{
			value: "vpc",
		},
		VPN: ChildrenJobInfoNetType{
			value: "vpn",
		},
		EIP: ChildrenJobInfoNetType{
			value: "eip",
		},
	}
}

func (c ChildrenJobInfoNetType) Value() string {
	return c.value
}

func (c ChildrenJobInfoNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoNetType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoTaskType struct {
	value string
}

type ChildrenJobInfoTaskTypeEnum struct {
	FULL_TRANS      ChildrenJobInfoTaskType
	FULL_INCR_TRANS ChildrenJobInfoTaskType
	INCR_TRANS      ChildrenJobInfoTaskType
}

func GetChildrenJobInfoTaskTypeEnum() ChildrenJobInfoTaskTypeEnum {
	return ChildrenJobInfoTaskTypeEnum{
		FULL_TRANS: ChildrenJobInfoTaskType{
			value: "FULL_TRANS 全量",
		},
		FULL_INCR_TRANS: ChildrenJobInfoTaskType{
			value: "FULL_INCR_TRANS 全量+增量",
		},
		INCR_TRANS: ChildrenJobInfoTaskType{
			value: "INCR_TRANS 增量",
		},
	}
}

func (c ChildrenJobInfoTaskType) Value() string {
	return c.value
}

func (c ChildrenJobInfoTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoTaskType) UnmarshalJSON(b []byte) error {
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
