package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteUserRepositoryAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteUserRepositoryAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserRepositoryAuthResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserRepositoryAuthResponse", string(data)}, " ")
}
