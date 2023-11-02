package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProjectRequest Request Object
type DeleteProjectRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 非核心项目删除立即删除标记
	XDeleteNow *bool `json:"X-Delete-Now,omitempty"`
}

func (o DeleteProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProjectRequest struct{}"
	}

	return strings.Join([]string{"DeleteProjectRequest", string(data)}, " ")
}
