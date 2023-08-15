package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertsRequest Request Object
type ListAlertsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	Body *DataobjectSearch `json:"body,omitempty"`
}

func (o ListAlertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertsRequest struct{}"
	}

	return strings.Join([]string{"ListAlertsRequest", string(data)}, " ")
}
