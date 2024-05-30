package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SyncStatusEnum 同步状态。 枚举值：   - NO_NEED: 未同步   - CREATE_SUCCESS: 创建成功   - CREATE_FAILED: 创建失败   - UPDATE_SUCCESS: 更新成功   - UPDATE_FAILED: 更新失败   - SUMMARY_SUCCESS: 整体成功   - SUMMARY_FAILED: 整体失败   - RUNNING: 整体运行   - OFFLINE: 已下线
type SyncStatusEnum struct {
	value string
}

type SyncStatusEnumEnum struct {
	NO_NEED         SyncStatusEnum
	CREATE_SUCCESS  SyncStatusEnum
	CREATE_FAILED   SyncStatusEnum
	UPDATE_SUCCESS  SyncStatusEnum
	UPDATE_FAILED   SyncStatusEnum
	SUMMARY_SUCCESS SyncStatusEnum
	SUMMARY_FAILED  SyncStatusEnum
	RUNNING         SyncStatusEnum
	OFFLINE         SyncStatusEnum
}

func GetSyncStatusEnumEnum() SyncStatusEnumEnum {
	return SyncStatusEnumEnum{
		NO_NEED: SyncStatusEnum{
			value: "NO_NEED",
		},
		CREATE_SUCCESS: SyncStatusEnum{
			value: "CREATE_SUCCESS",
		},
		CREATE_FAILED: SyncStatusEnum{
			value: "CREATE_FAILED",
		},
		UPDATE_SUCCESS: SyncStatusEnum{
			value: "UPDATE_SUCCESS",
		},
		UPDATE_FAILED: SyncStatusEnum{
			value: "UPDATE_FAILED",
		},
		SUMMARY_SUCCESS: SyncStatusEnum{
			value: "SUMMARY_SUCCESS",
		},
		SUMMARY_FAILED: SyncStatusEnum{
			value: "SUMMARY_FAILED",
		},
		RUNNING: SyncStatusEnum{
			value: "RUNNING",
		},
		OFFLINE: SyncStatusEnum{
			value: "OFFLINE",
		},
	}
}

func (c SyncStatusEnum) Value() string {
	return c.value
}

func (c SyncStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncStatusEnum) UnmarshalJSON(b []byte) error {
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
