package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 故事点
type IssueDetailResponseV4StoryPoint struct {

	// 故事点id
	Id *int32 `json:"id,omitempty"`

	// 故事点名称
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV4StoryPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV4StoryPoint struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV4StoryPoint", string(data)}, " ")
}
