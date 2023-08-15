package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProgramItemRequestBase 节目素材信息 。
type ProgramItemRequestBase struct {

	// 素材ID。
	MaterialId string `json:"materialId"`

	// 播放时长。
	PlayTime int32 `json:"playTime"`
}

func (o ProgramItemRequestBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgramItemRequestBase struct{}"
	}

	return strings.Join([]string{"ProgramItemRequestBase", string(data)}, " ")
}
