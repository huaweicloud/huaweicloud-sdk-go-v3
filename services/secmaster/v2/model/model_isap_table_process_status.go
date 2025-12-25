package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapTableProcessStatus **参数解释**: 处理状态 - COMPLETED 已完成 - CREATING 创建中 - UPDATING 更新中 - DELETING 删除中 - TRUNCATING 清空中 - UPGRADING   升级中       - CREATE_FAILED 创建失败 - UPDATE_FAILED 更新失败 - DELETE_FAILED 删除失败 - TRUNCATE_FAILED 清空失败 - UPGRADE_FAILED 升级失败        **约束限制** 不涉及 **取值范围**: - COMPLETED - CREATING - UPDATING - DELETING - TRUNCATING - UPGRADING       - CREATE_FAILED - UPDATE_FAILED - DELETE_FAILED - TRUNCATE_FAILED - UPGRADE_FAILED  **默认值** 不涉及
type IsapTableProcessStatus struct {
	value string
}

type IsapTableProcessStatusEnum struct {
	COMPLETED       IsapTableProcessStatus
	CREATING        IsapTableProcessStatus
	UPDATING        IsapTableProcessStatus
	DELETING        IsapTableProcessStatus
	TRUNCATING      IsapTableProcessStatus
	UPGRADING       IsapTableProcessStatus
	CREATE_FAILED   IsapTableProcessStatus
	UPDATE_FAILED   IsapTableProcessStatus
	DELETING_FAILED IsapTableProcessStatus
	TRUNCATE_FAILED IsapTableProcessStatus
	UPGRADE_FAILED  IsapTableProcessStatus
}

func GetIsapTableProcessStatusEnum() IsapTableProcessStatusEnum {
	return IsapTableProcessStatusEnum{
		COMPLETED: IsapTableProcessStatus{
			value: "COMPLETED",
		},
		CREATING: IsapTableProcessStatus{
			value: "CREATING",
		},
		UPDATING: IsapTableProcessStatus{
			value: "UPDATING",
		},
		DELETING: IsapTableProcessStatus{
			value: "DELETING",
		},
		TRUNCATING: IsapTableProcessStatus{
			value: "TRUNCATING",
		},
		UPGRADING: IsapTableProcessStatus{
			value: "UPGRADING",
		},
		CREATE_FAILED: IsapTableProcessStatus{
			value: "CREATE_FAILED",
		},
		UPDATE_FAILED: IsapTableProcessStatus{
			value: "UPDATE_FAILED",
		},
		DELETING_FAILED: IsapTableProcessStatus{
			value: "DELETING_FAILED",
		},
		TRUNCATE_FAILED: IsapTableProcessStatus{
			value: "TRUNCATE_FAILED",
		},
		UPGRADE_FAILED: IsapTableProcessStatus{
			value: "UPGRADE_FAILED",
		},
	}
}

func (c IsapTableProcessStatus) Value() string {
	return c.value
}

func (c IsapTableProcessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableProcessStatus) UnmarshalJSON(b []byte) error {
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
