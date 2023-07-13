package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserRepositoryAuthResponse Response Object
type UpdateUserRepositoryAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateUserRepositoryAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRepositoryAuthResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserRepositoryAuthResponse", string(data)}, " ")
}
