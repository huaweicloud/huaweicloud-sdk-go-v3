package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteWorkSpaceRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o DeleteWorkSpaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkSpaceRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkSpaceRequest", string(data)}, " ")
}
