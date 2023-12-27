package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShootScriptItemBaseInfo struct {

	// 剧本序号。
	SequenceNo *int32 `json:"sequence_no,omitempty"`
}

func (o ShootScriptItemBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptItemBaseInfo struct{}"
	}

	return strings.Join([]string{"ShootScriptItemBaseInfo", string(data)}, " ")
}
