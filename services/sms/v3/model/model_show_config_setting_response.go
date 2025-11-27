package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowConfigSettingResponse Response Object
type ShowConfigSettingResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 迁移类型 WINDOWS_MIGRATE_BLOCK: windows块级迁移 LINUX_MIGRATE_FILE: linux文件级迁移 LINUX_MIGRATE_BLOCK: linux块级迁移
	MigrateType *ShowConfigSettingResponseMigrateType `json:"migrate_type,omitempty"`

	// 配置项的具体配置信息
	Configurations *[]ConfigBody `json:"configurations,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowConfigSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigSettingResponse", string(data)}, " ")
}

type ShowConfigSettingResponseMigrateType struct {
	value string
}

type ShowConfigSettingResponseMigrateTypeEnum struct {
	WINDOWS_MIGRATE_BLOCK ShowConfigSettingResponseMigrateType
	LINUX_MIGRATE_FILE    ShowConfigSettingResponseMigrateType
	LINUX_MIGRATE_BLOCK   ShowConfigSettingResponseMigrateType
}

func GetShowConfigSettingResponseMigrateTypeEnum() ShowConfigSettingResponseMigrateTypeEnum {
	return ShowConfigSettingResponseMigrateTypeEnum{
		WINDOWS_MIGRATE_BLOCK: ShowConfigSettingResponseMigrateType{
			value: "WINDOWS_MIGRATE_BLOCK",
		},
		LINUX_MIGRATE_FILE: ShowConfigSettingResponseMigrateType{
			value: "LINUX_MIGRATE_FILE",
		},
		LINUX_MIGRATE_BLOCK: ShowConfigSettingResponseMigrateType{
			value: "LINUX_MIGRATE_BLOCK",
		},
	}
}

func (c ShowConfigSettingResponseMigrateType) Value() string {
	return c.value
}

func (c ShowConfigSettingResponseMigrateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigSettingResponseMigrateType) UnmarshalJSON(b []byte) error {
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
