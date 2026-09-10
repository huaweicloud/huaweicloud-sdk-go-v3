package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskLogsRequest Request Object
type ListTaskLogsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 作业任务ID。
	TaskId string `json:"task_id"`

	// 目录路径。
	Path *string `json:"path,omitempty"`
}

func (o ListTaskLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskLogsRequest struct{}"
	}

	return strings.Join([]string{"ListTaskLogsRequest", string(data)}, " ")
}
