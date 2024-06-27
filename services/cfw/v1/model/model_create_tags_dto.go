package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTagsDto struct {

	// 标签列表
	Tags *[]CreateTag `json:"tags,omitempty"`
}

func (o CreateTagsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsDto struct{}"
	}

	return strings.Join([]string{"CreateTagsDto", string(data)}, " ")
}
