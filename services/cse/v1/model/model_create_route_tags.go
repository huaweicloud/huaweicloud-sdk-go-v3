package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRouteTags 匹配的实例标记列表。
type CreateRouteTags struct {

	// 实例标记。满足标记条件的实例放到这一组。
	Tag *string `json:"&lt;tag&gt;,omitempty"`
}

func (o CreateRouteTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTags struct{}"
	}

	return strings.Join([]string{"CreateRouteTags", string(data)}, " ")
}
