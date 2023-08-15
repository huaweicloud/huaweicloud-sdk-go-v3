package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserResponse Response Object
type DeleteUserResponse struct {

	// 信息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserResponse", string(data)}, " ")
}
