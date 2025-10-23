package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LevelResourceCount 每类资源数量
type LevelResourceCount struct {

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源数量
	Count *int32 `json:"count,omitempty"`
}

func (o LevelResourceCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LevelResourceCount struct{}"
	}

	return strings.Join([]string{"LevelResourceCount", string(data)}, " ")
}
