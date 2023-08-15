package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerGroupsResponse Response Object
type DeleteServerGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupsResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupsResponse", string(data)}, " ")
}
