package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerGroupResponse Response Object
type DeleteServerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupResponse", string(data)}, " ")
}
