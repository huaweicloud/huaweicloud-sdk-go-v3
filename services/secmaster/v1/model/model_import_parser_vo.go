package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportParserVo struct {

	// 解析器文件名称
	FileName *string `json:"file_name,omitempty"`

	// 解析器名称
	ParserTitle *string `json:"parser_title,omitempty"`

	// 是否导入成功
	Success *bool `json:"success,omitempty"`
}

func (o ImportParserVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportParserVo struct{}"
	}

	return strings.Join([]string{"ImportParserVo", string(data)}, " ")
}
