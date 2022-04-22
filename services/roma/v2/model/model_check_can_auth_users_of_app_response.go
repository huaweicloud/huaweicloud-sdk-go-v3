package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckCanAuthUsersOfAppResponse struct {

	// 候选用户成员列表
	Users          *[]CandidatesUser `json:"users,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CheckCanAuthUsersOfAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCanAuthUsersOfAppResponse struct{}"
	}

	return strings.Join([]string{"CheckCanAuthUsersOfAppResponse", string(data)}, " ")
}
