package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentInstanceProjectId 中心网络附件对端实例的项目ID。
type AttachmentInstanceProjectId struct {

	// 中心网络附件对端实例的项目ID。
	AttachmentInstanceProjectId string `json:"attachment_instance_project_id"`
}

func (o AttachmentInstanceProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentInstanceProjectId struct{}"
	}

	return strings.Join([]string{"AttachmentInstanceProjectId", string(data)}, " ")
}
