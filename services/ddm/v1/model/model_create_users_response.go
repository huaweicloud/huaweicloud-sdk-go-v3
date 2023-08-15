package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUsersResponse Response Object
type CreateUsersResponse struct {

	// DDM实例帐号相关信息的集合。
	Users          *[]CreateUsersDetailResponses `json:"users,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o CreateUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUsersResponse struct{}"
	}

	return strings.Join([]string{"CreateUsersResponse", string(data)}, " ")
}
