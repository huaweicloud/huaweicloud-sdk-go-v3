package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTag struct {

	// 标签的key
	Key string `json:"key"`

	// 标签的value
	Value string `json:"value"`

	// 是否启用日志流标签
	TagsToStreamsEnable string `json:"tags_to_streams_enable"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
