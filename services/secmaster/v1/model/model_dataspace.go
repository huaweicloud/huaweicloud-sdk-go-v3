package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Dataspace struct {

	// 创建者
	CreateBy string `json:"create_by"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 工作空间ID
	DataspaceId string `json:"dataspace_id"`

	// 工作空间名称
	DataspaceName string `json:"dataspace_name"`

	// 数据空间类型；可选值：system-defined(系统预定义)、user-defined(用户自定义)
	DataspaceType *DataspaceDataspaceType `json:"dataspace_type,omitempty"`

	// 描述
	Description string `json:"description"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// region ID
	RegionId string `json:"region_id"`

	// 更新者
	UpdateBy string `json:"update_by"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

func (o Dataspace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dataspace struct{}"
	}

	return strings.Join([]string{"Dataspace", string(data)}, " ")
}

type DataspaceDataspaceType struct {
	value string
}

type DataspaceDataspaceTypeEnum struct {
	SYSTEM_DEFINED DataspaceDataspaceType
	USER_DEFINED   DataspaceDataspaceType
}

func GetDataspaceDataspaceTypeEnum() DataspaceDataspaceTypeEnum {
	return DataspaceDataspaceTypeEnum{
		SYSTEM_DEFINED: DataspaceDataspaceType{
			value: "system-defined",
		},
		USER_DEFINED: DataspaceDataspaceType{
			value: "user-defined",
		},
	}
}

func (c DataspaceDataspaceType) Value() string {
	return c.value
}

func (c DataspaceDataspaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataspaceDataspaceType) UnmarshalJSON(b []byte) error {
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
