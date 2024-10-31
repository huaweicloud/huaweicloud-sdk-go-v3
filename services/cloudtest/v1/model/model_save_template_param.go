package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaveTemplateParam struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 脑图ID
	MindmapId *string `json:"mindmap_id,omitempty"`

	// 脑图名称
	Name *string `json:"name,omitempty"`
}

func (o SaveTemplateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTemplateParam struct{}"
	}

	return strings.Join([]string{"SaveTemplateParam", string(data)}, " ")
}
