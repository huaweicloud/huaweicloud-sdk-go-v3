package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTagsV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 标签列表

	Responses      *[]string `json:"responses,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTagsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsV2Response struct{}"
	}

	return strings.Join([]string{"ListTagsV2Response", string(data)}, " ")
}
