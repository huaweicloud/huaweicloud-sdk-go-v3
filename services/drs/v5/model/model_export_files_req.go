package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportFilesReq struct {

	// 需要下载的文件名称。
	Files []string `json:"files"`
}

func (o ExportFilesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFilesReq struct{}"
	}

	return strings.Join([]string{"ExportFilesReq", string(data)}, " ")
}
