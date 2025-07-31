package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebTamperPrivilegedProcessRequestInfo struct {

	// 特权进程路径集合
	PrivilegedProcessPathList *[]string `json:"privileged_process_path_list,omitempty"`

	// 特权进程子进程可信状态
	PrivilegedChildStatus *bool `json:"privileged_child_status,omitempty"`
}

func (o WebTamperPrivilegedProcessRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperPrivilegedProcessRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperPrivilegedProcessRequestInfo", string(data)}, " ")
}
