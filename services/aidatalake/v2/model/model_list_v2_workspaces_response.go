package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2WorkspacesResponse Response Object
type ListV2WorkspacesResponse struct {

	// **参数解释**：总记录数。 **取值范围**：不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：工作空间列表，每个元素为一个Workspace对象，包含工作空间ID、名称、描述、创建时间等信息。
	Workspaces *[]Workspace `json:"workspaces,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListV2WorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2WorkspacesResponse struct{}"
	}

	return strings.Join([]string{"ListV2WorkspacesResponse", string(data)}, " ")
}
