package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserRequest Request Object
type ListUserRequest struct {
}

func (o ListUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserRequest struct{}"
	}

	return strings.Join([]string{"ListUserRequest", string(data)}, " ")
}
