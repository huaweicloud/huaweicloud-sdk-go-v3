package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTaskInfo 测试套执行结果
type ExecuteTaskInfo struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 发布版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 结果Code
	ResultCode *string `json:"result_code,omitempty"`

	// 任务执行结果Uri
	TestResultUri *string `json:"test_result_uri,omitempty"`

	// 状态Code
	StatusCode *string `json:"status_code,omitempty"`

	// 分支/迭代uri
	VersionUri *string `json:"version_uri,omitempty"`
}

func (o ExecuteTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTaskInfo struct{}"
	}

	return strings.Join([]string{"ExecuteTaskInfo", string(data)}, " ")
}
