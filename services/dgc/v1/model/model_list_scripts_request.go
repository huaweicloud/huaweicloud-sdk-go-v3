package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptsRequest Request Object
type ListScriptsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`
}

func (o ListScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptsRequest", string(data)}, " ")
}
