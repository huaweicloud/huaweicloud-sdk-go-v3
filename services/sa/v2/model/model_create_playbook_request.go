package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookRequest Request Object
type CreatePlaybookRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	Body *CreatePlaybookInfo `json:"body,omitempty"`
}

func (o CreatePlaybookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookRequest struct{}"
	}

	return strings.Join([]string{"CreatePlaybookRequest", string(data)}, " ")
}
