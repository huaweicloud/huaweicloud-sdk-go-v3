package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTagsReqbody struct {

	// 添加标签方式
	Action string `json:"action"`

	// 是否对外接口调用
	IsOpen bool `json:"is_open"`

	// 标签字段信息
	Tags []TagsBody `json:"tags"`
}

func (o CreateTagsReqbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsReqbody struct{}"
	}

	return strings.Join([]string{"CreateTagsReqbody", string(data)}, " ")
}
