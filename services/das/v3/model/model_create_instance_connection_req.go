package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateInstanceConnectionReq struct {

	// 数据库引擎类型。
	EngineType string `json:"engine_type"`

	// 网络类型。
	NetworkType CreateInstanceConnectionReqNetworkType `json:"network_type"`

	// 用户名。
	Username string `json:"username"`

	// 是否保存密码。
	IsSavePassword bool `json:"is_save_password"`

	// 密码。
	Password string `json:"password"`

	// 节点编号。
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 备注。
	Remarks *string `json:"remarks,omitempty"`

	// 端口。
	Port *int32 `json:"port,omitempty"`

	// 数据库名字。
	DatabaseName *string `json:"database_name,omitempty"`

	// sql记录开关。
	SqlRecordFlag *bool `json:"sql_record_flag,omitempty"`
}

func (o CreateInstanceConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceConnectionReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceConnectionReq", string(data)}, " ")
}

type CreateInstanceConnectionReqNetworkType struct {
	value string
}

type CreateInstanceConnectionReqNetworkTypeEnum struct {
	GAUSSDB CreateInstanceConnectionReqNetworkType
	DDS     CreateInstanceConnectionReqNetworkType
	RDS     CreateInstanceConnectionReqNetworkType
	ECS     CreateInstanceConnectionReqNetworkType
	GEMINI  CreateInstanceConnectionReqNetworkType
	DDM     CreateInstanceConnectionReqNetworkType
}

func GetCreateInstanceConnectionReqNetworkTypeEnum() CreateInstanceConnectionReqNetworkTypeEnum {
	return CreateInstanceConnectionReqNetworkTypeEnum{
		GAUSSDB: CreateInstanceConnectionReqNetworkType{
			value: "gaussdb",
		},
		DDS: CreateInstanceConnectionReqNetworkType{
			value: "dds",
		},
		RDS: CreateInstanceConnectionReqNetworkType{
			value: "rds",
		},
		ECS: CreateInstanceConnectionReqNetworkType{
			value: "ecs",
		},
		GEMINI: CreateInstanceConnectionReqNetworkType{
			value: "gemini",
		},
		DDM: CreateInstanceConnectionReqNetworkType{
			value: "ddm",
		},
	}
}

func (c CreateInstanceConnectionReqNetworkType) Value() string {
	return c.value
}

func (c CreateInstanceConnectionReqNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceConnectionReqNetworkType) UnmarshalJSON(b []byte) error {
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
