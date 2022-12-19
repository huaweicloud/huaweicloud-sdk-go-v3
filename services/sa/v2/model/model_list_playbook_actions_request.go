package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPlaybookActionsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`

	// request limit size
	Limit *int32 `json:"limit,omitempty"`

	// request offset, from 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPlaybookActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookActionsRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookActionsRequest", string(data)}, " ")
}
