package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
