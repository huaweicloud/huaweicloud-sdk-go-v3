package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIamGroupUsersRequest Request Object
type ListIamGroupUsersRequest struct {

	// IAM用户组id
	GroupId string `json:"group_id"`
}

func (o ListIamGroupUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIamGroupUsersRequest struct{}"
	}

	return strings.Join([]string{"ListIamGroupUsersRequest", string(data)}, " ")
}
