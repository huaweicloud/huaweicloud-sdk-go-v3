package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShootScriptItem 剧本参数配置。
type ShootScriptItem struct {

	// **参数解释**： 剧本序号。 **约束限制**： 同一个剧本序号不重复。 **默认取值**： 不涉及。
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	ShootScript *ShootScript `json:"shoot_script,omitempty"`
}

func (o ShootScriptItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptItem struct{}"
	}

	return strings.Join([]string{"ShootScriptItem", string(data)}, " ")
}
