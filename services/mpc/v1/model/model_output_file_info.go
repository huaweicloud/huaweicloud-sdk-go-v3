package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OutputFileInfo struct {

	// 输出文件名。
	OutputFileName *string `json:"output_file_name,omitempty" xml:"output_file_name"`

	// 处理信息。
	ExecDescription *string `json:"exec_description,omitempty" xml:"exec_description"`

	MetaData *SourceInfo `json:"meta_data,omitempty" xml:"meta_data"`
}

func (o OutputFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputFileInfo struct{}"
	}

	return strings.Join([]string{"OutputFileInfo", string(data)}, " ")
}
