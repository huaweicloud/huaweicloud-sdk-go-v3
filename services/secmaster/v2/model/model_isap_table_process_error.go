package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapTableProcessError **参数解释**: 表处理错误 - NONE 无 - MISSING_ASSOCIATIONS 关联缺失 - FAILED_INIT_STORAGE_RESOURCES_WHEN_CREATING 创建时初始化存储资源失败 - FAILED_INIT_FLINK_RESOURCES_WHEN_CREATING 创建时初始化 Flink 资源失败 - FAILED_DELETE_FLINK_RESOURCES_WHEN_DELETING 删除时删除 Flink 资源失败 - FAILED_DELETE_STORAGE_RESOURCES_WHEN_DELETING 删除时删除存储资源失败 - FAILED_DELETE_ALL_RESOURCES_WHEN_DELETING 删除时删除所有资源失败 - FAILED_UPDATE_STORAGE_SETTING 更新存储配置失败 - FAILED_UPDATE_FLINK_SCHEMA 更新 Flink 模式失败 - FAILED_UPDATE_STORAGE_SCHEMA 更新存储模式失败 - FAILED_TO_APPLY_RESOURCE 应用资源失败 - FAILED_TO_UPGRADE_RESOURCE_MODEL 升级资源模型失败  **约束限制** 不涉及 **取值范围**: - NONE - MISSING_ASSOCIATIONS - FAILED_INIT_STORAGE_RESOURCES_WHEN_CREATING - FAILED_INIT_FLINK_RESOURCES_WHEN_CREATING - FAILED_DELETE_FLINK_RESOURCES_WHEN_DELETING - FAILED_DELETE_STORAGE_RESOURCES_WHEN_DELETING - FAILED_DELETE_ALL_RESOURCES_WHEN_DELETING - FAILED_UPDATE_STORAGE_SETTING - FAILED_UPDATE_FLINK_SCHEMA - FAILED_UPDATE_STORAGE_SCHEMA - FAILED_TO_APPLY_RESOURCE - FAILED_TO_UPGRADE_RESOURCE_MODEL  **默认值** 不涉及
type IsapTableProcessError struct {
	value string
}

type IsapTableProcessErrorEnum struct {
	NONE                                          IsapTableProcessError
	MISSING_ASSOCIATIONS                          IsapTableProcessError
	FAILED_INIT_STORAGE_RESOURCES_WHEN_CREATING   IsapTableProcessError
	FAILED_INIT_FLINK_RESOURCES_WHEN_CREATING     IsapTableProcessError
	FAILED_DELETE_FLINK_RESOURCES_WHEN_DELETING   IsapTableProcessError
	FAILED_DELETE_STORAGE_RESOURCES_WHEN_DELETING IsapTableProcessError
	FAILED_DELETE_ALL_RESOURCES_WHEN_DELETING     IsapTableProcessError
	FAILED_UPDATE_STORAGE_SETTING                 IsapTableProcessError
	FAILED_UPDATE_FLINK_SCHEMA                    IsapTableProcessError
	FAILED_UPDATE_STORAGE_SCHEMA                  IsapTableProcessError
	FAILED_TO_APPLY_RESOURCE                      IsapTableProcessError
	FAILED_TO_UPGRADE_RESOURCE_MODEL              IsapTableProcessError
}

func GetIsapTableProcessErrorEnum() IsapTableProcessErrorEnum {
	return IsapTableProcessErrorEnum{
		NONE: IsapTableProcessError{
			value: "NONE",
		},
		MISSING_ASSOCIATIONS: IsapTableProcessError{
			value: "MISSING_ASSOCIATIONS",
		},
		FAILED_INIT_STORAGE_RESOURCES_WHEN_CREATING: IsapTableProcessError{
			value: "FAILED_INIT_STORAGE_RESOURCES_WHEN_CREATING",
		},
		FAILED_INIT_FLINK_RESOURCES_WHEN_CREATING: IsapTableProcessError{
			value: "FAILED_INIT_FLINK_RESOURCES_WHEN_CREATING",
		},
		FAILED_DELETE_FLINK_RESOURCES_WHEN_DELETING: IsapTableProcessError{
			value: "FAILED_DELETE_FLINK_RESOURCES_WHEN_DELETING",
		},
		FAILED_DELETE_STORAGE_RESOURCES_WHEN_DELETING: IsapTableProcessError{
			value: "FAILED_DELETE_STORAGE_RESOURCES_WHEN_DELETING",
		},
		FAILED_DELETE_ALL_RESOURCES_WHEN_DELETING: IsapTableProcessError{
			value: "FAILED_DELETE_ALL_RESOURCES_WHEN_DELETING",
		},
		FAILED_UPDATE_STORAGE_SETTING: IsapTableProcessError{
			value: "FAILED_UPDATE_STORAGE_SETTING",
		},
		FAILED_UPDATE_FLINK_SCHEMA: IsapTableProcessError{
			value: "FAILED_UPDATE_FLINK_SCHEMA",
		},
		FAILED_UPDATE_STORAGE_SCHEMA: IsapTableProcessError{
			value: "FAILED_UPDATE_STORAGE_SCHEMA",
		},
		FAILED_TO_APPLY_RESOURCE: IsapTableProcessError{
			value: "FAILED_TO_APPLY_RESOURCE",
		},
		FAILED_TO_UPGRADE_RESOURCE_MODEL: IsapTableProcessError{
			value: "FAILED_TO_UPGRADE_RESOURCE_MODEL",
		},
	}
}

func (c IsapTableProcessError) Value() string {
	return c.value
}

func (c IsapTableProcessError) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableProcessError) UnmarshalJSON(b []byte) error {
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
