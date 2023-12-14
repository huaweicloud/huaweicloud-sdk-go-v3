package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConsistencyResult 一致性校验的结果
type ConsistencyResult struct {

	// 校验目录
	DirCheck string `json:"dir_check"`

	// 文件总数
	NumTotalFiles int32 `json:"num_total_files"`

	// 差异文件数量
	NumDifferentFiles int32 `json:"num_different_files"`

	// 目的端缺少文件数量
	NumTargetMissFiles int32 `json:"num_target_miss_files"`

	// 目的端多余文件数量
	NumTargetMoreFiles int32 `json:"num_target_more_files"`
}

func (o ConsistencyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsistencyResult struct{}"
	}

	return strings.Join([]string{"ConsistencyResult", string(data)}, " ")
}
