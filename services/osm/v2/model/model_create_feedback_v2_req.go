package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFeedbackV2Req struct {

	// 举报类型
	Type string `json:"type"`

	// 涉案的华为云产品或服务; 举报网址
	Title string `json:"title"`

	// 侵权详情说明或举报内容及说明
	Content string `json:"content"`

	// 相关证明附件列表
	Files []string `json:"files"`

	Ip *IntellectualPropertyV2 `json:"ip,omitempty"`

	Contacts *ContactWayInfoV2 `json:"contacts"`
}

func (o CreateFeedbackV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFeedbackV2Req struct{}"
	}

	return strings.Join([]string{"CreateFeedbackV2Req", string(data)}, " ")
}
