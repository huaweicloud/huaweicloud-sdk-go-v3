package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Memory 内存信息。
type Memory struct {

	// 内存大小。
	Size *int32 `json:"size,omitempty"`

	// 内存单元数。
	Unit *string `json:"unit,omitempty"`
}

func (o Memory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Memory struct{}"
	}

	return strings.Join([]string{"Memory", string(data)}, " ")
}
