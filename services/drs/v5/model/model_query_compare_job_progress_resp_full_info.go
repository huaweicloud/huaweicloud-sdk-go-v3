package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryCompareJobProgressRespFullInfo 行对比与内容对比会返回该字段，全量对比信息。
type QueryCompareJobProgressRespFullInfo struct {

	// 全量数据对比进度，单位为%。
	Progress *float32 `json:"progress,omitempty"`

	// 全量数据对比速率。
	SrcSpeed *string `json:"src_speed,omitempty"`

	// 差异待复查行数。
	RecheckEntities *int32 `json:"recheck_entities,omitempty"`
}

func (o QueryCompareJobProgressRespFullInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCompareJobProgressRespFullInfo struct{}"
	}

	return strings.Join([]string{"QueryCompareJobProgressRespFullInfo", string(data)}, " ")
}
