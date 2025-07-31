package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryCompareJobProgressRespIncreInfo 动态内容对比会返回该字段，增量对比信息。
type QueryCompareJobProgressRespIncreInfo struct {

	// 增量对比时延，若时延为0表示所有增量数据都已对比完成。
	Delay *float32 `json:"delay,omitempty"`

	// 增量数据对比速率。
	SrcSpeed *string `json:"src_speed,omitempty"`

	// 每秒对比行数。
	Rps *int32 `json:"rps,omitempty"`

	// 增量位点。
	LogPoint *string `json:"log_point,omitempty"`

	// 差异待复查行数。
	RecheckEntities *int32 `json:"recheck_entities,omitempty"`
}

func (o QueryCompareJobProgressRespIncreInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCompareJobProgressRespIncreInfo struct{}"
	}

	return strings.Join([]string{"QueryCompareJobProgressRespIncreInfo", string(data)}, " ")
}
