package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMemberGroupResponse Response Object
type DeleteMemberGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMemberGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteMemberGroupResponse", string(data)}, " ")
}
