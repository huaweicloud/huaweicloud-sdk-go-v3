package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskInfo 删除测试套件参数
type DeleteTaskInfo struct {

	// 分支/迭代id
	VersionUri *string `json:"version_uri,omitempty"`

	// 任务id数组
	TaskUris *[]string `json:"task_uris,omitempty"`
}

func (o DeleteTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskInfo struct{}"
	}

	return strings.Join([]string{"DeleteTaskInfo", string(data)}, " ")
}
