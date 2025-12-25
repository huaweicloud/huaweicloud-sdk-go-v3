package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Pipe struct {

	// 创建者
	CreateBy string `json:"create_by"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`

	// 数据空间名称
	DataspaceName string `json:"dataspace_name"`

	// 描述
	Description string `json:"description"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 数据管道ID
	PipeId string `json:"pipe_id"`

	// 数据管道名称
	PipeName string `json:"pipe_name"`

	// 数据管道类型；可选值：system-defined(系统预定义)、user-defined(用户自定义)
	PipeType *PipePipeType `json:"pipe_type,omitempty"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 分区个数；默认创建1个，最大支持创建64个分区
	Shards int32 `json:"shards"`

	// 数据的保存时间，单位为天；默认30天，取值范围为1~3600
	StoragePeriod int32 `json:"storage_period"`

	// 更新者
	UpdateBy string `json:"update_by"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

func (o Pipe) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pipe struct{}"
	}

	return strings.Join([]string{"Pipe", string(data)}, " ")
}

type PipePipeType struct {
	value string
}

type PipePipeTypeEnum struct {
	SYSTEM_DEFINED PipePipeType
	USER_DEFINED   PipePipeType
}

func GetPipePipeTypeEnum() PipePipeTypeEnum {
	return PipePipeTypeEnum{
		SYSTEM_DEFINED: PipePipeType{
			value: "system-defined",
		},
		USER_DEFINED: PipePipeType{
			value: "user-defined",
		},
	}
}

func (c PipePipeType) Value() string {
	return c.value
}

func (c PipePipeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipePipeType) UnmarshalJSON(b []byte) error {
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
