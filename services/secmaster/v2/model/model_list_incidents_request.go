package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIncidentsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	Body *DataobjectSearch `json:"body,omitempty"`
}

func (o ListIncidentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncidentsRequest struct{}"
	}

	return strings.Join([]string{"ListIncidentsRequest", string(data)}, " ")
}
