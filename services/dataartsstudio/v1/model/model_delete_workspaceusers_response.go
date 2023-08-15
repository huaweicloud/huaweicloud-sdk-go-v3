package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkspaceusersResponse Response Object
type DeleteWorkspaceusersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWorkspaceusersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkspaceusersResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkspaceusersResponse", string(data)}, " ")
}
