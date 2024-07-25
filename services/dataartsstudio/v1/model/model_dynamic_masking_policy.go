package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DynamicMaskingPolicy struct {

	// 字段脱敏策略id。
	Id *string `json:"id,omitempty"`

	// 动态脱敏策略id。
	PolicySetId *string `json:"policy_set_id,omitempty"`

	// 数据表中的字段名称。
	ColumnName *string `json:"column_name,omitempty"`

	// 数据表中字段的数据类型。
	ColumnType *string `json:"column_type,omitempty"`

	// 具体动态脱敏规则参数介绍请参见[动态脱敏规则介绍](dataartsstudio_01_1036.html)。 HIVE数据源动态脱敏算法 - MASK 掩盖英文字符和数字 - MASK_SHOW_LAST_4 保留后四位 - MASK_SHOW_FIRST_4 保留前四位 - MASK_HASH 哈希掩盖 - MASK_DATE_SHOW_YEAR 掩盖月份和日期 - MASK_NULL NULL掩盖  DWS数据源动态脱敏算法 - DWS_ALL_MASK 全掩码 - DWS_BACK_KEEP 保留后4位，其余脱敏为* - DWS_FRONT_KEEP 保留前2位，其余脱敏为* - DWS_SELF_CONFIG 需要输入开始位置、结束位置、脱敏字符传入detail结构体中，例如{\"start\": 1, \"end\": 2, \"string_target\": \"*\"}  [DLI数据源动态脱敏算法](tag:nohcs) - [MASK 掩盖英文字符和数字](tag:nohcs) - [MASK_SHOW_LAST_4 保留后四位](tag:nohcs) - [MASK_SHOW_FIRST_4 保留前四位](tag:nohcs) - [MASK_HASH 哈希掩盖](tag:nohcs) - [MASK_DATE_SHOW_YEAR 掩盖月份和日期](tag:nohcs) - [MASK_NULL NULL掩盖](tag:nohcs)
	AlgorithmType *string `json:"algorithm_type,omitempty"`

	// 同步状态： - UNKNOWN 未知状态 - NOT_SYNC 未同步 - SYNCING 同步中 - SYNC_SUCCESS 同步成功 - SYNC_FAIL 同步失败 - SYNC_PARTIAL_FAIL 存在失败 - DELETE_FAIL 删除失败 - DELETING 删除中 - UPDATING 更新中 - DATA_UPDATED 数据存在更新
	SyncStatus *DynamicMaskingPolicySyncStatus `json:"sync_status,omitempty"`

	// 动态脱敏策略算法详情。
	AlgorithmDetail *string `json:"algorithm_detail,omitempty"`

	AlgorithmDetailDto *AlgorithmDetailDto `json:"algorithm_detail_dto,omitempty"`
}

func (o DynamicMaskingPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DynamicMaskingPolicy struct{}"
	}

	return strings.Join([]string{"DynamicMaskingPolicy", string(data)}, " ")
}

type DynamicMaskingPolicySyncStatus struct {
	value string
}

type DynamicMaskingPolicySyncStatusEnum struct {
	UNKNOWN           DynamicMaskingPolicySyncStatus
	NOT_SYNC          DynamicMaskingPolicySyncStatus
	SYNCING           DynamicMaskingPolicySyncStatus
	SYNC_SUCCESS      DynamicMaskingPolicySyncStatus
	SYNC_FAIL         DynamicMaskingPolicySyncStatus
	SYNC_PARTIAL_FAIL DynamicMaskingPolicySyncStatus
	DELETE_FAIL       DynamicMaskingPolicySyncStatus
	DELETING          DynamicMaskingPolicySyncStatus
	UPDATING          DynamicMaskingPolicySyncStatus
	DATA_UPDATED      DynamicMaskingPolicySyncStatus
}

func GetDynamicMaskingPolicySyncStatusEnum() DynamicMaskingPolicySyncStatusEnum {
	return DynamicMaskingPolicySyncStatusEnum{
		UNKNOWN: DynamicMaskingPolicySyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: DynamicMaskingPolicySyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: DynamicMaskingPolicySyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: DynamicMaskingPolicySyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: DynamicMaskingPolicySyncStatus{
			value: "SYNC_FAIL",
		},
		SYNC_PARTIAL_FAIL: DynamicMaskingPolicySyncStatus{
			value: "SYNC_PARTIAL_FAIL",
		},
		DELETE_FAIL: DynamicMaskingPolicySyncStatus{
			value: "DELETE_FAIL",
		},
		DELETING: DynamicMaskingPolicySyncStatus{
			value: "DELETING",
		},
		UPDATING: DynamicMaskingPolicySyncStatus{
			value: "UPDATING",
		},
		DATA_UPDATED: DynamicMaskingPolicySyncStatus{
			value: "DATA_UPDATED",
		},
	}
}

func (c DynamicMaskingPolicySyncStatus) Value() string {
	return c.value
}

func (c DynamicMaskingPolicySyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DynamicMaskingPolicySyncStatus) UnmarshalJSON(b []byte) error {
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
