package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupResponse Response Object
type UpdateGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateGroupResponse", string(data)}, " ")
}
