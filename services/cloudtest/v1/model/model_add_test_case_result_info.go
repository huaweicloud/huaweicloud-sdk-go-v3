package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddTestCaseResultInfo struct {

	// 主键
	Uri *string `json:"uri,omitempty"`

	// 结果名字
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 用例结果
	Result *string `json:"result,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// 前置条件
	Preparation *string `json:"preparation,omitempty"`

	// 用例步骤结果信息
	Steps *[]TestCaseStepResultInfo `json:"steps,omitempty"`

	// 版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 任务URI
	TaskUri *string `json:"task_uri,omitempty"`

	// 测试套结果URI
	TaskResultUri *string `json:"task_result_uri,omitempty"`
}

func (o AddTestCaseResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTestCaseResultInfo struct{}"
	}

	return strings.Join([]string{"AddTestCaseResultInfo", string(data)}, " ")
}
