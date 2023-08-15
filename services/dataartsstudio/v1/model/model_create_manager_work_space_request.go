package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManagerWorkSpaceRequest Request Object
type CreateManagerWorkSpaceRequest struct {

	// DataArtsStudio实例id
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
