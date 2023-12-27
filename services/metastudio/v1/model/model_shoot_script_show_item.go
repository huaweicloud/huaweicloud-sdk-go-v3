package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShootScriptShowItem 剧本参数配置。
type ShootScriptShowItem struct {

	// 剧本序号。
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	ShootScript *ShootScriptDetail `json:"shoot_script,omitempty"`
}

func (o ShootScriptShowItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptShowItem struct{}"
	}

	return strings.Join([]string{"ShootScriptShowItem", string(data)}, " ")
}
