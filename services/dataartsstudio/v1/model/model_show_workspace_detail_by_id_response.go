package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceDetailByIdResponse Response Object
type ShowWorkspaceDetailByIdResponse struct {
	Data           *CreateWorkspaceResultData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowWorkspaceDetailByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceDetailByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceDetailByIdResponse", string(data)}, " ")
}
