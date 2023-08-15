package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkSpaceResponse Response Object
type CreateWorkSpaceResponse struct {

	// 工作空间id
	WorkspaceId    *string `json:"workspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkSpaceResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkSpaceResponse", string(data)}, " ")
}
