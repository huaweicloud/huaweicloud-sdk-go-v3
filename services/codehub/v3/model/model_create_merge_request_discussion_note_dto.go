package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMergeRequestDiscussionNoteDto struct {

	// 检视意见内容
	Body string `json:"body"`
}

func (o CreateMergeRequestDiscussionNoteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestDiscussionNoteDto struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestDiscussionNoteDto", string(data)}, " ")
}
