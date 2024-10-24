package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFabricWorkspaceTagsResponse Response Object
type ListFabricWorkspaceTagsResponse struct {

	// 资源标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 系统标签列表。
	SysTags *[]SystemTag `json:"sys_tags,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFabricWorkspaceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFabricWorkspaceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListFabricWorkspaceTagsResponse", string(data)}, " ")
}
