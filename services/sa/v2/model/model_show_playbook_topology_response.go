package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookTopologyResponse Response Object
type ShowPlaybookTopologyResponse struct {

	// tatal count
	Count *int32 `json:"count,omitempty"`

	// Playbook action instances list
	ActionInstances *[]ActionInstanceInfo `json:"action_instances,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPlaybookTopologyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookTopologyResponse struct{}"
	}

	return strings.Join([]string{"ShowPlaybookTopologyResponse", string(data)}, " ")
}
