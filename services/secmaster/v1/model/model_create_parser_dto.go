package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateParserDto struct {

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 相关描述信息
	Modules *[]CreateModuleVo `json:"modules,omitempty"`

	// UUID
	ParserId *string `json:"parser_id,omitempty"`

	// 相关标题
	Title string `json:"title"`
}

func (o CreateParserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateParserDto struct{}"
	}

	return strings.Join([]string{"CreateParserDto", string(data)}, " ")
}
