package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookInstanceRequest Request Object
type ShowPlaybookInstanceRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// instance _id
	InstanceId string `json:"instance_id"`
}

func (o ShowPlaybookInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookInstanceRequest", string(data)}, " ")
}
