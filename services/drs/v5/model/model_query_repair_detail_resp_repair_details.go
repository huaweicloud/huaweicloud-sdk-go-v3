package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRepairDetailRespRepairDetails struct {

	// 源表标志列值。
	SourceMeta *string `json:"source_meta,omitempty"`

	// 目标表标志列值。
	TargetMeta *string `json:"target_meta,omitempty"`

	// 修复SQL状态。
	RepairSqlState *int32 `json:"repair_sql_state,omitempty"`

	// 修复SQL。
	RepairSqlInfo *string `json:"repair_sql_info,omitempty"`
}

func (o QueryRepairDetailRespRepairDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRepairDetailRespRepairDetails struct{}"
	}

	return strings.Join([]string{"QueryRepairDetailRespRepairDetails", string(data)}, " ")
}
