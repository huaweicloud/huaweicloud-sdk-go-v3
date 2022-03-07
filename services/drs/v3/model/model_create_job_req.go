package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建任务请求体
type CreateJobReq struct {
	// 是否绑定eip，网络类型为eip时必填且为true

	BindEip *bool `json:"bind_eip,omitempty"`
	// 迁移场景，migration-实时迁移,sync-实时同步,cloudDataGuard-实时灾备

	DbUseType CreateJobReqDbUseType `json:"db_use_type"`
	// 任务名称，约束：任务名称在4位到50位之间，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。

	Name string `json:"name"`
	// 任务描述。  **约束**：任务描述不能超过256位，且不能包含!<>&'\"\\特殊字符。

	Description *string `json:"description,omitempty"`
	// 引擎类型 - mysql：迁移，同步使用 - mongodb：迁移使用 - cloudDataGuard-mysql：灾备使用 - gaussdbv5，postgresql：同步使用

	EngineType CreateJobReqEngineType `json:"engine_type"`
	// 指定目标实例是否限制为只读，MySQL迁移和灾备，且job_direction为up时设置有效。（灾备场景下，单主灾备且本云为备为必填且为true，不填默认设置为true）。

	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`
	// 迁移方向，up ：入云 ，灾备场景时对应本云为备，down：出云，灾备场景时对应本云为主，non-dbs：自建

	JobDirection CreateJobReqJobDirection `json:"job_direction"`
	// - db_use_type 是cloudDataGuard时，必填，灾备类型是双主灾备时 muti_write取值true, 否则为false。 - db_use_type 是其他类型时，muti_write是非必选参数

	MultiWrite *bool `json:"multi_write,omitempty"`
	// 网络类型

	NetType CreateJobReqNetType `json:"net_type"`
	// 节点个数。MongoDB数据库时对应源端分片个数，源库为集群时必填，[1-32]，MySQL双主灾备时会默认设置为2。

	NodeNum *int32 `json:"node_num,omitempty"`
	// 规格类型。

	NodeType CreateJobReqNodeType `json:"node_type"`

	SourceEndpoint *Endpoint `json:"source_endpoint"`

	TargetEndpoint *Endpoint `json:"target_endpoint"`
	// 标签信息。

	Tags *[]ResourceTag `json:"tags,omitempty"`
	// 迁移模式，FULL_TRANS 全量,FULL_INCR_TRANS 全量+增量,INCR_TRANS 增量，灾备场景单主灾备仅支持全量加增量（FULL_INCR_TRANS）

	TaskType CreateJobReqTaskType `json:"task_type"`
	// DRS实例所在子网ID，对应目标库相同VPC下已创建的子网（subnet）的网络ID，UUID格式。

	CustomizeSutnetId string `json:"customize_sutnet_id"`
	// 产品id。

	ProductId *string `json:"product_id,omitempty"`
	// 企业项目，不填默认为default，key值必须为_sys_enterprise_project_id，value为企业项目ID，只能有一个企业项目。

	SysTags *[]ResourceTag `json:"sys_tags,omitempty"`
	// 任务处于异常状态一段时间后，将会自动结束，单位为天。(范围14-100)，不传默认为14天。

	ExpiredDays *string `json:"expired_days,omitempty"`
}

func (o CreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobReq struct{}"
	}

	return strings.Join([]string{"CreateJobReq", string(data)}, " ")
}

type CreateJobReqDbUseType struct {
	value string
}

type CreateJobReqDbUseTypeEnum struct {
	MIGRATION        CreateJobReqDbUseType
	SYNC             CreateJobReqDbUseType
	CLOUD_DATA_GUARD CreateJobReqDbUseType
}

func GetCreateJobReqDbUseTypeEnum() CreateJobReqDbUseTypeEnum {
	return CreateJobReqDbUseTypeEnum{
		MIGRATION: CreateJobReqDbUseType{
			value: "migration",
		},
		SYNC: CreateJobReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: CreateJobReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c CreateJobReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqDbUseType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqEngineType struct {
	value string
}

type CreateJobReqEngineTypeEnum struct {
	MYSQL                  CreateJobReqEngineType
	MONGODB                CreateJobReqEngineType
	CLOUD_DATA_GUARD_MYSQL CreateJobReqEngineType
	GAUSSDBV5              CreateJobReqEngineType
	POSTGRESQL             CreateJobReqEngineType
}

func GetCreateJobReqEngineTypeEnum() CreateJobReqEngineTypeEnum {
	return CreateJobReqEngineTypeEnum{
		MYSQL: CreateJobReqEngineType{
			value: "mysql",
		},
		MONGODB: CreateJobReqEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: CreateJobReqEngineType{
			value: "cloudDataGuard-mysql",
		},
		GAUSSDBV5: CreateJobReqEngineType{
			value: "gaussdbv5",
		},
		POSTGRESQL: CreateJobReqEngineType{
			value: "postgresql",
		},
	}
}

func (c CreateJobReqEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqEngineType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqJobDirection struct {
	value string
}

type CreateJobReqJobDirectionEnum struct {
	UP      CreateJobReqJobDirection
	DOWN    CreateJobReqJobDirection
	NON_DBS CreateJobReqJobDirection
}

func GetCreateJobReqJobDirectionEnum() CreateJobReqJobDirectionEnum {
	return CreateJobReqJobDirectionEnum{
		UP: CreateJobReqJobDirection{
			value: "up",
		},
		DOWN: CreateJobReqJobDirection{
			value: "down",
		},
		NON_DBS: CreateJobReqJobDirection{
			value: "non-dbs",
		},
	}
}

func (c CreateJobReqJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqJobDirection) UnmarshalJSON(b []byte) error {
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

type CreateJobReqNetType struct {
	value string
}

type CreateJobReqNetTypeEnum struct {
	VPN CreateJobReqNetType
	VPC CreateJobReqNetType
	EIP CreateJobReqNetType
}

func GetCreateJobReqNetTypeEnum() CreateJobReqNetTypeEnum {
	return CreateJobReqNetTypeEnum{
		VPN: CreateJobReqNetType{
			value: "vpn",
		},
		VPC: CreateJobReqNetType{
			value: "vpc",
		},
		EIP: CreateJobReqNetType{
			value: "eip",
		},
	}
}

func (c CreateJobReqNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqNetType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqNodeType struct {
	value string
}

type CreateJobReqNodeTypeEnum struct {
	HIGH CreateJobReqNodeType
}

func GetCreateJobReqNodeTypeEnum() CreateJobReqNodeTypeEnum {
	return CreateJobReqNodeTypeEnum{
		HIGH: CreateJobReqNodeType{
			value: "high",
		},
	}
}

func (c CreateJobReqNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqNodeType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqTaskType struct {
	value string
}

type CreateJobReqTaskTypeEnum struct {
	FULL_TRANS      CreateJobReqTaskType
	FULL_INCR_TRANS CreateJobReqTaskType
	INCR_TRANS      CreateJobReqTaskType
}

func GetCreateJobReqTaskTypeEnum() CreateJobReqTaskTypeEnum {
	return CreateJobReqTaskTypeEnum{
		FULL_TRANS: CreateJobReqTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: CreateJobReqTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: CreateJobReqTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c CreateJobReqTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqTaskType) UnmarshalJSON(b []byte) error {
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
