package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelWorkspaceRequest Request Object
type CancelWorkspaceRequest struct {
}

func (o CancelWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"CancelWorkspaceRequest", string(data)}, " ")
}
