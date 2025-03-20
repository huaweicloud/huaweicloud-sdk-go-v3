package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceTagsV5Response Response Object
type ListResourceTagsV5Response struct {

	// 自定义标签列表。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceTagsV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTagsV5Response struct{}"
	}

	return strings.Join([]string{"ListResourceTagsV5Response", string(data)}, " ")
}
