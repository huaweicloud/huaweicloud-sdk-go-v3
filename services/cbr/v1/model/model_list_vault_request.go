package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListVaultRequest struct {
	// 每页显示条目数，正整数

	Limit *int32 `json:"limit,omitempty"`
	// 存储库名称

	Name *string `json:"name,omitempty"`
	// 偏移值,正整数

	Offset *int32 `json:"offset,omitempty"`
	// 云类型

	CloudType *ListVaultRequestCloudType `json:"cloud_type,omitempty"`
	// 保护类型

	ProtectType *ListVaultRequestProtectType `json:"protect_type,omitempty"`
	// 资源类型

	ObjectType *string `json:"object_type,omitempty"`
	// 企业项目id或all_granted_eps，all_granted_eps表示查询用户有权限的所有企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 存储库ID

	Id *string `json:"id,omitempty"`
	// 策略ID

	PolicyId *string `json:"policy_id,omitempty"`
	// 状态

	Status *string `json:"status,omitempty"`
	// 资源id

	ResourceIds *string `json:"resource_ids,omitempty"`
}

func (o ListVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultRequest struct{}"
	}

	return strings.Join([]string{"ListVaultRequest", string(data)}, " ")
}

type ListVaultRequestCloudType struct {
	value string
}

type ListVaultRequestCloudTypeEnum struct {
	PUBLIC ListVaultRequestCloudType
	HYBRID ListVaultRequestCloudType
}

func GetListVaultRequestCloudTypeEnum() ListVaultRequestCloudTypeEnum {
	return ListVaultRequestCloudTypeEnum{
		PUBLIC: ListVaultRequestCloudType{
			value: "public",
		},
		HYBRID: ListVaultRequestCloudType{
			value: "hybrid",
		},
	}
}

func (c ListVaultRequestCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVaultRequestCloudType) UnmarshalJSON(b []byte) error {
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

type ListVaultRequestProtectType struct {
	value string
}

type ListVaultRequestProtectTypeEnum struct {
	BACKUP      ListVaultRequestProtectType
	REPLICATION ListVaultRequestProtectType
}

func GetListVaultRequestProtectTypeEnum() ListVaultRequestProtectTypeEnum {
	return ListVaultRequestProtectTypeEnum{
		BACKUP: ListVaultRequestProtectType{
			value: "backup",
		},
		REPLICATION: ListVaultRequestProtectType{
			value: "replication",
		},
	}
}

func (c ListVaultRequestProtectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVaultRequestProtectType) UnmarshalJSON(b []byte) error {
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
