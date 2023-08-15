package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveUsersResponse Response Object
type RemoveUsersResponse struct {
	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RemoveUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveUsersResponse struct{}"
	}

	return strings.Join([]string{"RemoveUsersResponse", string(data)}, " ")
}
