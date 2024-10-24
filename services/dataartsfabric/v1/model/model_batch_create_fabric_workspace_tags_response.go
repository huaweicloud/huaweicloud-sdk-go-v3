package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateFabricWorkspaceTagsResponse Response Object
type BatchCreateFabricWorkspaceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateFabricWorkspaceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateFabricWorkspaceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateFabricWorkspaceTagsResponse", string(data)}, " ")
}
