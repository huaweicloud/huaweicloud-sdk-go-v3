package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMembersRequest struct {
}

func (o ListMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMembersRequest struct{}"
	}

	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}
