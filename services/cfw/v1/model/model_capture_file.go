package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaptureFile struct {

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 下载链接
	Url *string `json:"url,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`
}

func (o CaptureFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaptureFile struct{}"
	}

	return strings.Join([]string{"CaptureFile", string(data)}, " ")
}
