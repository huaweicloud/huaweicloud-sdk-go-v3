package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesByTagRequestBody list resource tags request
type ListResourcesByTagRequestBody struct {

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 匹配项列表。
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ListResourcesByTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagRequestBody struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagRequestBody", string(data)}, " ")
}
