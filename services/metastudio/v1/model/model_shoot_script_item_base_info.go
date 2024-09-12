package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShootScriptItemBaseInfo struct {

	// **参数解释**： 剧本序号。 **约束限制**： 同一个剧本序号不重复。 **默认取值**： 不涉及。
	SequenceNo *int32 `json:"sequence_no,omitempty"`
}

func (o ShootScriptItemBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptItemBaseInfo struct{}"
	}

	return strings.Join([]string{"ShootScriptItemBaseInfo", string(data)}, " ")
}
