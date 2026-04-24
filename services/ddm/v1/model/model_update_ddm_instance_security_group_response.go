package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDdmInstanceSecurityGroupResponse Response Object
type UpdateDdmInstanceSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDdmInstanceSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDdmInstanceSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateDdmInstanceSecurityGroupResponse", string(data)}, " ")
}
