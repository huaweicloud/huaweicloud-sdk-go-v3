package model

import (
	"encoding/json"

	"strings"
)

type TransTemplateGroup struct {
	// 模板组名称
	Name *string `json:"name,omitempty"`
	// 视频信息列表
	Videos      *[]VideoObj  `json:"videos,omitempty"`
	Audio       *Audio       `json:"audio,omitempty"`
	VideoCommon *VideoCommon `json:"video_common,omitempty"`
	Common      *Common      `json:"common,omitempty"`
}

func (o TransTemplateGroup) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TransTemplateGroup struct{}"
	}

	return strings.Join([]string{"TransTemplateGroup", string(data)}, " ")
}
