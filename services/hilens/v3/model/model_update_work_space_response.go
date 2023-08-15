package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkSpaceResponse Response Object
type UpdateWorkSpaceResponse struct {

	// 工作空间id
	WorkspaceId    *string `json:"workspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateWorkSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkSpaceResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkSpaceResponse", string(data)}, " ")
}
