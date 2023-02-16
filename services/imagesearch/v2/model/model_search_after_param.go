package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 最后一个结果的排序信息，用于全量召回场景。目前KEYWORD/FACE搜索和条件检查支持全量召回。
type SearchAfterParam struct {

	// 结果的得分。
	Score *float64 `json:"score,omitempty"`

	// 结果的唯一ID。
	Id *string `json:"id,omitempty"`
}

func (o SearchAfterParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAfterParam struct{}"
	}

	return strings.Join([]string{"SearchAfterParam", string(data)}, " ")
}
