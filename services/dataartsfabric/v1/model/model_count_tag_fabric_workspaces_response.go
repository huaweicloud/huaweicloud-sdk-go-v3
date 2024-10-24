package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTagFabricWorkspacesResponse Response Object
type CountTagFabricWorkspacesResponse struct {

	// 实例总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountTagFabricWorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTagFabricWorkspacesResponse struct{}"
	}

	return strings.Join([]string{"CountTagFabricWorkspacesResponse", string(data)}, " ")
}
