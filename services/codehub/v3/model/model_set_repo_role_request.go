package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetRepoRoleRequest struct {
	// 用户id

	MemberId string `json:"member_id"`
	// 仓库uuid

	RepositoryUuid string `json:"repository_uuid"`

	Body *SetRepoRoleRequestBody `json:"body,omitempty"`
}

func (o SetRepoRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRepoRoleRequest struct{}"
	}

	return strings.Join([]string{"SetRepoRoleRequest", string(data)}, " ")
}
