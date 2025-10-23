package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVaultRequest Request Object
type ListVaultRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 归属region
	RegionId *string `json:"region_id,omitempty"`

	// 企业项目id或all_granted_eps，all_granted_eps表示查询用户有权限的所有企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 合规id
	ComplianceId *string `json:"compliance_id,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 资源id，支持多资源，以英文逗号分隔
	ResourceIds *string `json:"resource_ids,omitempty"`

	// 状态，available可用，lock锁定，frozen冻结，error错误，deleting删除中
	Status *ListVaultRequestStatus `json:"status,omitempty"`

	// 存储库存储类型，server云服务器，disk磁盘，turbo，workspace云空间，vmware，rds，file
	ObjectType *ListVaultRequestObjectType `json:"object_type,omitempty"`

	// 存储库保护类型，backup：备份，replication：复制
	ProtectType *ListVaultRequestProtectType `json:"protect_type,omitempty"`

	// 每页显示条目数，正整数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移值,正整数
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultRequest struct{}"
	}

	return strings.Join([]string{"ListVaultRequest", string(data)}, " ")
}

type ListVaultRequestStatus struct {
	value string
}

type ListVaultRequestStatusEnum struct {
	AVAILABLE ListVaultRequestStatus
	LOCK      ListVaultRequestStatus
	FROZEN    ListVaultRequestStatus
	ERROR     ListVaultRequestStatus
	DELETING  ListVaultRequestStatus
}

func GetListVaultRequestStatusEnum() ListVaultRequestStatusEnum {
	return ListVaultRequestStatusEnum{
		AVAILABLE: ListVaultRequestStatus{
			value: "available",
		},
		LOCK: ListVaultRequestStatus{
			value: "lock",
		},
		FROZEN: ListVaultRequestStatus{
			value: "frozen",
		},
		ERROR: ListVaultRequestStatus{
			value: "error",
		},
		DELETING: ListVaultRequestStatus{
			value: "deleting",
		},
	}
}

func (c ListVaultRequestStatus) Value() string {
	return c.value
}

func (c ListVaultRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVaultRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListVaultRequestObjectType struct {
	value string
}

type ListVaultRequestObjectTypeEnum struct {
	SERVER    ListVaultRequestObjectType
	DISK      ListVaultRequestObjectType
	TURBO     ListVaultRequestObjectType
	WORKSPACE ListVaultRequestObjectType
	VMWARE    ListVaultRequestObjectType
	RDS       ListVaultRequestObjectType
	FILE      ListVaultRequestObjectType
}

func GetListVaultRequestObjectTypeEnum() ListVaultRequestObjectTypeEnum {
	return ListVaultRequestObjectTypeEnum{
		SERVER: ListVaultRequestObjectType{
			value: "server",
		},
		DISK: ListVaultRequestObjectType{
			value: "disk",
		},
		TURBO: ListVaultRequestObjectType{
			value: "turbo",
		},
		WORKSPACE: ListVaultRequestObjectType{
			value: "workspace",
		},
		VMWARE: ListVaultRequestObjectType{
			value: "vmware",
		},
		RDS: ListVaultRequestObjectType{
			value: "rds",
		},
		FILE: ListVaultRequestObjectType{
			value: "file",
		},
	}
}

func (c ListVaultRequestObjectType) Value() string {
	return c.value
}

func (c ListVaultRequestObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVaultRequestObjectType) UnmarshalJSON(b []byte) error {
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

func (c ListVaultRequestProtectType) Value() string {
	return c.value
}

func (c ListVaultRequestProtectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVaultRequestProtectType) UnmarshalJSON(b []byte) error {
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
