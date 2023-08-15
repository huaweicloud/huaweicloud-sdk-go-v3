package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepoMemberRequest Request Object
type DeleteRepoMemberRequest struct {

	// 仓库成员id
	MemberId string `json:"member_id"`

	// 仓库uuid(由CreateRepository接口返回)
	RepositoryUuid string `json:"repository_uuid"`
}

func (o DeleteRepoMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepoMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteRepoMemberRequest", string(data)}, " ")
}
