package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPlaybookTopologyRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// instance _id
	InstanceId string `json:"instance_id"`
}

func (o ShowPlaybookTopologyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookTopologyRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookTopologyRequest", string(data)}, " ")
}
