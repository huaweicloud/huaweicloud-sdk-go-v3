package model

import (
	"encoding/json"

	"strings"
)

// 转码生成文件信息。仅当转码成功后才能查询到此信息，未转码、正在转码以及转码失败时，此字段为空。
type TranscodeInfo struct {
	// 转码模板组名称

	TemplateGroupName string `json:"template_group_name"`

	Output []Output `json:"output"`

	ExecDesc *string `json:"exec_desc,omitempty"`

	TranscodeStatus *string `json:"transcode_status,omitempty"`
}

func (o TranscodeInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TranscodeInfo struct{}"
	}

	return strings.Join([]string{"TranscodeInfo", string(data)}, " ")
}
