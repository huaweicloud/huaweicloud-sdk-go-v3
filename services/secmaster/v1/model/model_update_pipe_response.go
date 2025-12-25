package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePipeResponse Response Object
type UpdatePipeResponse struct {

	// 创建者
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 数据空间名称
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 数据管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 数据管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 数据管道类型；可选值：system-defined(系统预定义)、user-defined(用户自定义)
	PipeType *UpdatePipeResponsePipeType `json:"pipe_type,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 分区个数；默认创建1个，最大支持创建64个分区
	Shards *int32 `json:"shards,omitempty"`

	// 数据的保存时间，单位为天；默认30天，取值范围为1~3600
	StoragePeriod *int32 `json:"storage_period,omitempty"`

	// 更新者
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdatePipeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipeResponse struct{}"
	}

	return strings.Join([]string{"UpdatePipeResponse", string(data)}, " ")
}

type UpdatePipeResponsePipeType struct {
	value string
}

type UpdatePipeResponsePipeTypeEnum struct {
	SYSTEM_DEFINED UpdatePipeResponsePipeType
	USER_DEFINED   UpdatePipeResponsePipeType
}

func GetUpdatePipeResponsePipeTypeEnum() UpdatePipeResponsePipeTypeEnum {
	return UpdatePipeResponsePipeTypeEnum{
		SYSTEM_DEFINED: UpdatePipeResponsePipeType{
			value: "system-defined",
		},
		USER_DEFINED: UpdatePipeResponsePipeType{
			value: "user-defined",
		},
	}
}

func (c UpdatePipeResponsePipeType) Value() string {
	return c.value
}

func (c UpdatePipeResponsePipeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePipeResponsePipeType) UnmarshalJSON(b []byte) error {
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
