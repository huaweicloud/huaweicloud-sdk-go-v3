package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AomMappingfilesInfo struct {
	// 路径名

	FileName string `json:"file_name"`

	LogStreamInfo *AomMappingLogStreamInfo `json:"log_stream_info"`
}

func (o AomMappingfilesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AomMappingfilesInfo struct{}"
	}

	return strings.Join([]string{"AomMappingfilesInfo", string(data)}, " ")
}
