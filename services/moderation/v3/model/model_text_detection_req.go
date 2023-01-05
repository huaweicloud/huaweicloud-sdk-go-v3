package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 文本内容审核请求体
type TextDetectionReq struct {

	// 事件类型。  可选值如下：  nickname: 昵称  title: 标题  article: 帖⼦  comment: 评论  barrage: 弹幕  search: 搜索栏  profile: 个⼈简介
	EventType string `json:"event_type"`

	// 检测时使用的自定义黑名单词库列表。
	GlossaryNames *[]string `json:"glossary_names,omitempty"`

	Data *TextDetectionDataReq `json:"data"`

	// 检测时使用的自定义白名单词库列表。
	WhiteGlossaryNames *[]string `json:"white_glossary_names,omitempty"`
}

func (o TextDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionReq struct{}"
	}

	return strings.Join([]string{"TextDetectionReq", string(data)}, " ")
}
