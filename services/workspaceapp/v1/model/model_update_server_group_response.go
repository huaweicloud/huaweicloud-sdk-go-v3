package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerGroupResponse Response Object
type UpdateServerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerGroupResponse", string(data)}, " ")
}
