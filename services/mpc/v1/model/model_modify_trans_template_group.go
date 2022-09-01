package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyTransTemplateGroup struct {

	// 模板组ID
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 模板组名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 视频信息列表
	Videos *[]VideoObj `json:"videos,omitempty" xml:"videos"`

	Audio *Audio `json:"audio,omitempty" xml:"audio"`

	VideoCommon *VideoCommon `json:"video_common,omitempty" xml:"video_common"`

	Common *Common `json:"common,omitempty" xml:"common"`
}

func (o ModifyTransTemplateGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTransTemplateGroup struct{}"
	}

	return strings.Join([]string{"ModifyTransTemplateGroup", string(data)}, " ")
}
