package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDataspaceResponse Response Object
type ShowDataspaceResponse struct {

	// 创建者
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 工作空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 工作空间名称
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 数据空间类型；可选值：system-defined(系统预定义)、user-defined(用户自定义)
	DataspaceType *ShowDataspaceResponseDataspaceType `json:"dataspace_type,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// region ID
	RegionId *string `json:"region_id,omitempty"`

	// 更新者
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDataspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataspaceResponse struct{}"
	}

	return strings.Join([]string{"ShowDataspaceResponse", string(data)}, " ")
}

type ShowDataspaceResponseDataspaceType struct {
	value string
}

type ShowDataspaceResponseDataspaceTypeEnum struct {
	SYSTEM_DEFINED ShowDataspaceResponseDataspaceType
	USER_DEFINED   ShowDataspaceResponseDataspaceType
}

func GetShowDataspaceResponseDataspaceTypeEnum() ShowDataspaceResponseDataspaceTypeEnum {
	return ShowDataspaceResponseDataspaceTypeEnum{
		SYSTEM_DEFINED: ShowDataspaceResponseDataspaceType{
			value: "system-defined",
		},
		USER_DEFINED: ShowDataspaceResponseDataspaceType{
			value: "user-defined",
		},
	}
}

func (c ShowDataspaceResponseDataspaceType) Value() string {
	return c.value
}

func (c ShowDataspaceResponseDataspaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataspaceResponseDataspaceType) UnmarshalJSON(b []byte) error {
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
