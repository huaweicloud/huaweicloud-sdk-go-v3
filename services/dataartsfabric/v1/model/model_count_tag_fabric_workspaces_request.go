package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTagFabricWorkspacesRequest Request Object
type CountTagFabricWorkspacesRequest struct {
	Body *ListTagWorkspacesRequestBody `json:"body,omitempty"`
}

func (o CountTagFabricWorkspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTagFabricWorkspacesRequest struct{}"
	}

	return strings.Join([]string{"CountTagFabricWorkspacesRequest", string(data)}, " ")
}
