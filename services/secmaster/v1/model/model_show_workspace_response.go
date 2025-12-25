package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceResponse Response Object
type ShowWorkspaceResponse struct {
	Workspace      *ShowWorkspaceResponseBodyWorkspace `json:"workspace,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceResponse", string(data)}, " ")
}
