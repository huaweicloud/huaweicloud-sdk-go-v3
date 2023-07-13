package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIamGroupsRequest Request Object
type ListIamGroupsRequest struct {
}

func (o ListIamGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIamGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListIamGroupsRequest", string(data)}, " ")
}
