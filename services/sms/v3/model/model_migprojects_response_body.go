package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 批量查询迁移项目返回的迁移项目信息
type MigprojectsResponseBody struct {

	// 迁移项目ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 迁移项目名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 是否使用公网IP迁移
	UsePublicIp *bool `json:"use_public_ip,omitempty" xml:"use_public_ip"`

	// 是否为默认模板
	Isdefault *bool `json:"isdefault,omitempty" xml:"isdefault"`

	// 迁移后是否启动目的端虚拟机
	StartTargetServer *bool `json:"start_target_server,omitempty" xml:"start_target_server"`

	// 区域名称
	Region *string `json:"region,omitempty" xml:"region"`

	// 模板中配置的限速信息，单位：Mbps
	SpeedLimit *int32 `json:"speed_limit,omitempty" xml:"speed_limit"`

	// 迁移项目下是否存在服务器
	ExistServer *bool `json:"exist_server,omitempty" xml:"exist_server"`

	// 迁移项目描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 迁移项目默认迁移类型
	Type *MigprojectsResponseBodyType `json:"type,omitempty" xml:"type"`

	// 迁移项目所属的企业项目名称
	EnterpriseProject *string `json:"enterprise_project,omitempty" xml:"enterprise_project"`

	// 是否持续同步
	Syncing *bool `json:"syncing,omitempty" xml:"syncing"`
}

func (o MigprojectsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigprojectsResponseBody struct{}"
	}

	return strings.Join([]string{"MigprojectsResponseBody", string(data)}, " ")
}

type MigprojectsResponseBodyType struct {
	value string
}

type MigprojectsResponseBodyTypeEnum struct {
	MIGRATE_BLOCK MigprojectsResponseBodyType
	MIGRATE_FILE  MigprojectsResponseBodyType
}

func GetMigprojectsResponseBodyTypeEnum() MigprojectsResponseBodyTypeEnum {
	return MigprojectsResponseBodyTypeEnum{
		MIGRATE_BLOCK: MigprojectsResponseBodyType{
			value: "MIGRATE_BLOCK",
		},
		MIGRATE_FILE: MigprojectsResponseBodyType{
			value: "MIGRATE_FILE",
		},
	}
}

func (c MigprojectsResponseBodyType) Value() string {
	return c.value
}

func (c MigprojectsResponseBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigprojectsResponseBodyType) UnmarshalJSON(b []byte) error {
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
