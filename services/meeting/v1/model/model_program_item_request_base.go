package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 新增/更新节目素材信息请求
type ProgramItemRequestBase struct {
	// 素材ID

	MaterialId string `json:"materialId"`
	// 播放时长

	PlayTime int32 `json:"playTime"`
}

func (o ProgramItemRequestBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgramItemRequestBase struct{}"
	}

	return strings.Join([]string{"ProgramItemRequestBase", string(data)}, " ")
}
