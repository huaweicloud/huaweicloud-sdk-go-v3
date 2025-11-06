package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PermissionsDto struct {
	RepositoryAccess *MemberAccess `json:"repository_access,omitempty"`

	GroupAccess *MemberAccess `json:"group_access,omitempty"`
}

func (o PermissionsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionsDto struct{}"
	}

	return strings.Join([]string{"PermissionsDto", string(data)}, " ")
}
