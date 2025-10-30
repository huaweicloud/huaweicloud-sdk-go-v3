package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateHostsGroupResponse Response Object
type AssociateHostsGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateHostsGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateHostsGroupResponse struct{}"
	}

	return strings.Join([]string{"AssociateHostsGroupResponse", string(data)}, " ")
}
