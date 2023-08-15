package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkSpaceResponse Response Object
type DeleteWorkSpaceResponse struct {

	// 工作空间id
	WorkspaceId    *string `json:"workspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteWorkSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkSpaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkSpaceResponse", string(data)}, " ")
}
