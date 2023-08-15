package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePlaybookInstanceRequest Request Object
type ChangePlaybookInstanceRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// instance _id
	InstanceId string `json:"instance_id"`

	Body *OperationPlaybookInfo `json:"body,omitempty"`
}

func (o ChangePlaybookInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePlaybookInstanceRequest struct{}"
	}

	return strings.Join([]string{"ChangePlaybookInstanceRequest", string(data)}, " ")
}
