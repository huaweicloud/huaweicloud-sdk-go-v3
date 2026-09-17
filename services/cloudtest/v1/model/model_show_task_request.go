package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskRequest Request Object
type ShowTaskRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 测试套件uri
	TaskUri string `json:"task_uri"`

	// 分支/迭代uri
	VersionUri *string `json:"version_uri,omitempty"`
}

func (o ShowTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRequest", string(data)}, " ")
}
