package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManagerWorkSpaceRequest Request Object
type CreateManagerWorkSpaceRequest struct {

	// DataArts Studio实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	InstanceId string `json:"instance_id"`

	Body *CreateWorkspaceParams `json:"body,omitempty"`
}

func (o CreateManagerWorkSpaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManagerWorkSpaceRequest struct{}"
	}

	return strings.Join([]string{"CreateManagerWorkSpaceRequest", string(data)}, " ")
}
