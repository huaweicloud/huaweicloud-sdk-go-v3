package model

import (
	"encoding/json"

	"strings"
)

type OutputFileInfo struct {
	// 输出文件名。

	OutputFileName *string `json:"output_file_name,omitempty"`
	// 处理信息。

	ExecDescription *string `json:"exec_description,omitempty"`

	MetaData *MetaData `json:"meta_data,omitempty"`
}

func (o OutputFileInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OutputFileInfo struct{}"
	}

	return strings.Join([]string{"OutputFileInfo", string(data)}, " ")
}
