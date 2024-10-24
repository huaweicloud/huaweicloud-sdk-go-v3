package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagFabricWorkspacesResponse Response Object
type ListTagFabricWorkspacesResponse struct {

	// 资源列表。
	Resources *[]TagFabricWorkspace `json:"resources,omitempty"`

	// 资源总数
	TotalCount *int32 `json:"total_count,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTagFabricWorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagFabricWorkspacesResponse struct{}"
	}

	return strings.Join([]string{"ListTagFabricWorkspacesResponse", string(data)}, " ")
}
