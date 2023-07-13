package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIpsProtectModeUsingPostResponse Response Object
type ChangeIpsProtectModeUsingPostResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ChangeIpsProtectModeUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIpsProtectModeUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ChangeIpsProtectModeUsingPostResponse", string(data)}, " ")
}
