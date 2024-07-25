package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DynamicMaskingPolicyUpdate struct {

	// 数据表中的字段名称。
	ColumnName string `json:"column_name"`

	// 数据表中字段的数据类型。
	ColumnType string `json:"column_type"`

	// 具体动态脱敏规则参数介绍请参见[动态脱敏规则介绍](dataartsstudio_01_1036.html)。 HIVE数据源动态脱敏算法 - MASK 掩盖英文字符和数字 - MASK_SHOW_LAST_4 保留后四位 - MASK_SHOW_FIRST_4 保留前四位 - MASK_HASH 哈希掩盖 - MASK_DATE_SHOW_YEAR 掩盖月份和日期 - MASK_NULL NULL掩盖  DWS数据源动态脱敏算法 - DWS_ALL_MASK 全掩码 - DWS_BACK_KEEP 保留后4位，其余脱敏为* - DWS_FRONT_KEEP 保留前2位，其余脱敏为* - DWS_SELF_CONFIG 需要输入开始位置、结束位置、脱敏字符传入detail结构体中，例如{\"start\": 1, \"end\": 2, \"string_target\": \"*\"}  [DLI数据源动态脱敏算法](tag:nohcs) - [MASK 掩盖英文字符和数字](tag:nohcs) - [MASK_SHOW_LAST_4 保留后四位](tag:nohcs) - [MASK_SHOW_FIRST_4 保留前四位](tag:nohcs) - [MASK_HASH 哈希掩盖](tag:nohcs) - [MASK_DATE_SHOW_YEAR 掩盖月份和日期](tag:nohcs) - [MASK_NULL NULL掩盖](tag:nohcs)
	AlgorithmType *string `json:"algorithm_type,omitempty"`

	// 动态脱敏策略算法详情。
	AlgorithmDetail *string `json:"algorithm_detail,omitempty"`

	AlgorithmDetailDto *AlgorithmDetailDto `json:"algorithm_detail_dto,omitempty"`
}

func (o DynamicMaskingPolicyUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DynamicMaskingPolicyUpdate struct{}"
	}

	return strings.Join([]string{"DynamicMaskingPolicyUpdate", string(data)}, " ")
}
