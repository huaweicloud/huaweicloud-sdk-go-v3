package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestCaseCommentInfo struct {

	// 评论
	Comment string `json:"comment"`

	// 通知人列表
	Notifier *[]string `json:"notifier,omitempty"`

	// 分支uri/测试计划uri
	VersionUri *string `json:"version_uri,omitempty"`
}

func (o TestCaseCommentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseCommentInfo struct{}"
	}

	return strings.Join([]string{"TestCaseCommentInfo", string(data)}, " ")
}
