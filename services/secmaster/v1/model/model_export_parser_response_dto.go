package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportParserResponseDto struct {

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 相关描述信息
	Modules *[]ShowModuleExportVo `json:"modules,omitempty"`

	// 相关标题
	Title *string `json:"title,omitempty"`
}

func (o ExportParserResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportParserResponseDto struct{}"
	}

	return strings.Join([]string{"ExportParserResponseDto", string(data)}, " ")
}
