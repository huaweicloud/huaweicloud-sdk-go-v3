package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteUserGroupsResponse Response Object
type BatchDeleteUserGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteUserGroupsResponse", string(data)}, " ")
}
