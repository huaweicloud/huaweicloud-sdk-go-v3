package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFabricWorkspaceTagsResponse Response Object
type BatchDeleteFabricWorkspaceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteFabricWorkspaceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFabricWorkspaceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteFabricWorkspaceTagsResponse", string(data)}, " ")
}
