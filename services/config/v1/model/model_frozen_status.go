package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrozenStatus 资源记录器冻结状态
type FrozenStatus struct {

	// 是否冻结
	IsFrozen *bool `json:"is_frozen,omitempty"`

	// 冻结场景
	FrozenScene *[]string `json:"frozen_scene,omitempty"`
}

func (o FrozenStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrozenStatus struct{}"
	}

	return strings.Join([]string{"FrozenStatus", string(data)}, " ")
}
