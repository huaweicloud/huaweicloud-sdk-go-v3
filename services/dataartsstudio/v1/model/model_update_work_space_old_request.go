package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkSpaceOldRequest Request Object
type UpdateWorkSpaceOldRequest struct {

	// DataArts Studio实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	InstanceId string `json:"instance_id"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *WorkspaceDto `json:"body,omitempty"`
}

func (o UpdateWorkSpaceOldRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkSpaceOldRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkSpaceOldRequest", string(data)}, " ")
}
