package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServerGroupResponse Response Object
type CreateServerGroupResponse struct {
	ServerGroup    *CreateServerGroupResult `json:"server_group,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateServerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateServerGroupResponse", string(data)}, " ")
}
