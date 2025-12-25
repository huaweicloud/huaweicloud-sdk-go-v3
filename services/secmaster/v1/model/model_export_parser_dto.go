package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportParserDto struct {

	// 相关描述信息
	ParserIds []string `json:"parser_ids"`
}

func (o ExportParserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportParserDto struct{}"
	}

	return strings.Join([]string{"ExportParserDto", string(data)}, " ")
}
