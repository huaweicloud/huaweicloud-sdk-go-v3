package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCaseRequest Request Object
type ShowTestCaseRequest struct {

	// 用例id
	TestcaseId string `json:"testcase_id"`

	// 分支uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 任务
	TaskUri *string `json:"task_uri,omitempty"`

	// 是否刷新缓存
	Refresh *bool `json:"refresh,omitempty"`

	// 是否回收站资源
	IsRecycle *bool `json:"is_recycle,omitempty"`
}

func (o ShowTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseRequest struct{}"
	}

	return strings.Join([]string{"ShowTestCaseRequest", string(data)}, " ")
}
