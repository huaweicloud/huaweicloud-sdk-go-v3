package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserRepositoryAuthResponse Response Object
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
