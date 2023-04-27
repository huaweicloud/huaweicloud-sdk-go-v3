package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIamUsersRequest struct {
}

func (o ListIamUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIamUsersRequest struct{}"
	}

	return strings.Join([]string{"ListIamUsersRequest", string(data)}, " ")
}
