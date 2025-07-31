package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivilegedProcessResponseInfo struct {

	// 特权进程文件路径
	ProcessFilePath *string `json:"process_file_path,omitempty"`
}

func (o PrivilegedProcessResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivilegedProcessResponseInfo struct{}"
	}

	return strings.Join([]string{"PrivilegedProcessResponseInfo", string(data)}, " ")
}
