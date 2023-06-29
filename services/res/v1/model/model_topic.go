package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Topic
type Topic struct {

	// 名称。
	Name *string `json:"name,omitempty"`

	// id。
	Id *string `json:"id,omitempty"`
}

func (o Topic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Topic struct{}"
	}

	return strings.Join([]string{"Topic", string(data)}, " ")
}
