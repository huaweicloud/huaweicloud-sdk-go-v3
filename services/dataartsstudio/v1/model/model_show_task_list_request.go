package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskListRequest Request Object
type ShowTaskListRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *QueryTaskRequest `json:"body,omitempty"`
}

func (o ShowTaskListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskListRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskListRequest", string(data)}, " ")
}
