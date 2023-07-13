package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkspacesRequest Request Object
type DeleteWorkspacesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id数组
	Ids []int64 `json:"ids"`
}

func (o DeleteWorkspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkspacesRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkspacesRequest", string(data)}, " ")
}
