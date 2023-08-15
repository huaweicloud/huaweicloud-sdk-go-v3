package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueItemSfV4StoryPoint id 值， 1 '0'， 2 '1/2'， 3 '1' ， 4 '2' ， 5 '3' ， 6 '5' ， 7 '8' ， 8 '13' ， 9 '21' ， 10 '40' ， 11 '80' ， 12 '100' ， 13 '∞' ， 14 '?' ，
type IssueItemSfV4StoryPoint struct {

	// 故事点id
	Id *int32 `json:"id,omitempty"`

	// 故事点
	Name *string `json:"name,omitempty"`
}

func (o IssueItemSfV4StoryPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4StoryPoint struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4StoryPoint", string(data)}, " ")
}
