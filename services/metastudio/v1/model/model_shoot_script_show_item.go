package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShootScriptShowItem 剧本参数配置。
type ShootScriptShowItem struct {

	// **参数解释**： 剧本序号。 **约束限制**： 同一个剧本序号不重复。 **默认取值**： 不涉及。
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	ShootScript *ShootScriptDetail `json:"shoot_script,omitempty"`

	SubtitleFileInfo *SubtitleFileInfo `json:"subtitle_file_info,omitempty"`
}

func (o ShootScriptShowItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptShowItem struct{}"
	}

	return strings.Join([]string{"ShootScriptShowItem", string(data)}, " ")
}
