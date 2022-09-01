package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransTemplateGroup struct {

	// 模板组名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 视频信息列表
	Videos *[]VideoObj `json:"videos,omitempty" xml:"videos"`

	Audio *Audio `json:"audio,omitempty" xml:"audio"`

	VideoCommon *VideoCommon `json:"video_common,omitempty" xml:"video_common"`

	Common *Common `json:"common,omitempty" xml:"common"`
}

func (o TransTemplateGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransTemplateGroup struct{}"
	}

	return strings.Join([]string{"TransTemplateGroup", string(data)}, " ")
}
