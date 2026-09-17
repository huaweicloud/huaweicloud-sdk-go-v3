package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestVersionCaseRequest Request Object
type ShowTestVersionCaseRequest struct {

	// 用例id
	CaseUri string `json:"case_uri"`

	// 分支uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 任务
	TaskUri *string `json:"taskUri,omitempty"`

	// 是否刷新缓存
	Refresh *bool `json:"refresh,omitempty"`

	// 是否回收站资源
	IsRecycle *bool `json:"is_recycle,omitempty"`
}

func (o ShowTestVersionCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestVersionCaseRequest struct{}"
	}

	return strings.Join([]string{"ShowTestVersionCaseRequest", string(data)}, " ")
}
