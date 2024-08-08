package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnbindInstanceTags struct {

	// 实例标签
	Tags []Tag `json:"tags"`
}

func (o UnbindInstanceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindInstanceTags struct{}"
	}

	return strings.Join([]string{"UnbindInstanceTags", string(data)}, " ")
}
