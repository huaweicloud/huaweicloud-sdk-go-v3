package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupMembershipResponse Response Object
type DeleteGroupMembershipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGroupMembershipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupMembershipResponse struct{}"
	}

	return strings.Join([]string{"DeleteGroupMembershipResponse", string(data)}, " ")
}
