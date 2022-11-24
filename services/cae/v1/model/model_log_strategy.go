package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志策略。
type LogStrategy struct {

	// 收集路径。
	CollectPath *string `json:"collect_path,omitempty"`

	// 挂载路径。
	MountPath *string `json:"mount_path,omitempty"`
}

func (o LogStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogStrategy struct{}"
	}

	return strings.Join([]string{"LogStrategy", string(data)}, " ")
}
