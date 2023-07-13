package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIndicatorRequest Request Object
type UpdateIndicatorRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// workspace id
	WorkspaceId string `json:"workspace_id"`

	// ID of indicator
	IndicatorId string `json:"indicator_id"`

	Body *UpdateIndicatorRequestBody `json:"body,omitempty"`
}

func (o UpdateIndicatorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIndicatorRequest struct{}"
	}

	return strings.Join([]string{"UpdateIndicatorRequest", string(data)}, " ")
}
