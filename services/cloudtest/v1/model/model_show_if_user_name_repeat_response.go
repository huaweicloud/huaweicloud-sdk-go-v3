package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIfUserNameRepeatResponse Response Object
type ShowIfUserNameRepeatResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIfUserNameRepeatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIfUserNameRepeatResponse struct{}"
	}

	return strings.Join([]string{"ShowIfUserNameRepeatResponse", string(data)}, " ")
}
