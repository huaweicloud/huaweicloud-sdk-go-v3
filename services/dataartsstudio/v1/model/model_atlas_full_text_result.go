package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AtlasFullTextResult 请求参数
type AtlasFullTextResult struct {
	Entity *AtlasEntityHeader `json:"entity,omitempty"`

	// 数值
	Score float32 `json:"score,omitempty"`
}

func (o AtlasFullTextResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtlasFullTextResult struct{}"
	}

	return strings.Join([]string{"AtlasFullTextResult", string(data)}, " ")
}
