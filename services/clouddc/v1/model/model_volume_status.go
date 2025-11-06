package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeStatus 逻辑卷状态总览
type VolumeStatus struct {

	// **参数解释**： 逻辑卷状态 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	State *string `json:"state,omitempty"`

	Health *Health `json:"health,omitempty"`
}

func (o VolumeStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeStatus struct{}"
	}

	return strings.Join([]string{"VolumeStatus", string(data)}, " ")
}
