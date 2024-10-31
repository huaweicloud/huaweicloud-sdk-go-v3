package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeUserPrivilegeGroupRequest Request Object
type ChangeUserPrivilegeGroupRequest struct {
	Body *ChangeUserPrivilegeGroupReq `json:"body,omitempty"`
}

func (o ChangeUserPrivilegeGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeUserPrivilegeGroupRequest struct{}"
	}

	return strings.Join([]string{"ChangeUserPrivilegeGroupRequest", string(data)}, " ")
}
