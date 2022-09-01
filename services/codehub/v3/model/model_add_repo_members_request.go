package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddRepoMembersRequest struct {

	// 仓库uuid
	RepositoryUuid string `json:"repository_uuid" xml:"repository_uuid"`

	Body *CreateRepoMemberRequest `json:"body,omitempty" xml:"body"`
}

func (o AddRepoMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRepoMembersRequest struct{}"
	}

	return strings.Join([]string{"AddRepoMembersRequest", string(data)}, " ")
}
