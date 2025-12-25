package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableProcessStatus **参数解释**: 表处理状态 - COMPLETED 处理完成 - CREATING 创建中 - UPDATING 更新中 - DELETING 删除中 - TRUNCATING 清空中 - CREATE_FAILED 创建失败 - UPDATING_FAILED 更新失败 - DELETING_FAILED 删除失败 - TRUNCATE_FAILED 清空失败  **约束限制** 不涉及 **取值范围**: - COMPLETED - CREATING - UPDATING - DELETING - TRUNCATING - CREATE_FAILED - UPDATING_FAILED - DELETING_FAILED - TRUNCATE_FAILED  **默认值** 不涉及
type TableProcessStatus struct {
	value string
}

type TableProcessStatusEnum struct {
	COMPLETED       TableProcessStatus
	CREATING        TableProcessStatus
	UPDATING        TableProcessStatus
	DELETING        TableProcessStatus
	TRUNCATING      TableProcessStatus
	CREATE_FAILED   TableProcessStatus
	UPDATING_FAILED TableProcessStatus
	DELETING_FAILED TableProcessStatus
	TRUNCATE_FAILED TableProcessStatus
}

func GetTableProcessStatusEnum() TableProcessStatusEnum {
	return TableProcessStatusEnum{
		COMPLETED: TableProcessStatus{
			value: "COMPLETED",
		},
		CREATING: TableProcessStatus{
			value: "CREATING",
		},
		UPDATING: TableProcessStatus{
			value: "UPDATING",
		},
		DELETING: TableProcessStatus{
			value: "DELETING",
		},
		TRUNCATING: TableProcessStatus{
			value: "TRUNCATING",
		},
		CREATE_FAILED: TableProcessStatus{
			value: "CREATE_FAILED",
		},
		UPDATING_FAILED: TableProcessStatus{
			value: "UPDATING_FAILED",
		},
		DELETING_FAILED: TableProcessStatus{
			value: "DELETING_FAILED",
		},
		TRUNCATE_FAILED: TableProcessStatus{
			value: "TRUNCATE_FAILED",
		},
	}
}

func (c TableProcessStatus) Value() string {
	return c.value
}

func (c TableProcessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableProcessStatus) UnmarshalJSON(b []byte) error {
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
