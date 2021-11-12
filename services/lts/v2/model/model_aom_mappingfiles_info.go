package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AomMappingfilesInfo struct {
	// 路径名

	FileName string `json:"file_name"`
	// 接入规则详情

	LogStreamInfo *AomMappingLogStreamInfo `json:"log_stream_info"`
}

func (o AomMappingfilesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AomMappingfilesInfo struct{}"
	}

	return strings.Join([]string{"AomMappingfilesInfo", string(data)}, " ")
}
