package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBinlogParseRequestBody 查看binlog概览请求体
type ShowBinlogParseRequestBody struct {

	// binlog文件名称
	FileName string `json:"file_name"`

	// 页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ShowBinlogParseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogParseRequestBody struct{}"
	}

	return strings.Join([]string{"ShowBinlogParseRequestBody", string(data)}, " ")
}
