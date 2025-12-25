package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkspaceResponse Response Object
type DeleteWorkspaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkspaceResponse", string(data)}, " ")
}
