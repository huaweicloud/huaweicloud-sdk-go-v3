package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopProjectRequest Request Object
type UpdateTopProjectRequest struct {

	// 平台空间ID，您可以在平台单击所需的空间名称，进入空间设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *UpdateTopProjectReq `json:"body,omitempty"`
}

func (o UpdateTopProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopProjectRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopProjectRequest", string(data)}, " ")
}
