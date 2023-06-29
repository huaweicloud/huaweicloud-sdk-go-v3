package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddWorkSpaceUsersResponse Response Object
type AddWorkSpaceUsersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddWorkSpaceUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWorkSpaceUsersResponse struct{}"
	}

	return strings.Join([]string{"AddWorkSpaceUsersResponse", string(data)}, " ")
}
