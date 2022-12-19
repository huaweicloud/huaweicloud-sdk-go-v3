package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteIndicatorRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// workspace id
	WorkspaceId string `json:"workspace_id"`

	Body *DeleteIndicatorRequestBody `json:"body,omitempty"`
}

func (o DeleteIndicatorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIndicatorRequest struct{}"
	}

	return strings.Join([]string{"DeleteIndicatorRequest", string(data)}, " ")
}
