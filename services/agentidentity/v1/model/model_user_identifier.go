package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserIdentifier OAuth2.0 token or user ID used to generate the workload access token, only one field is required (either user_token or user_id)
type UserIdentifier struct {

	// OAuth2.0 token for user identification
	UserToken *string `json:"user_token,omitempty"`

	// User ID for identification
	UserId *string `json:"user_id,omitempty"`
}

func (o UserIdentifier) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserIdentifier struct{}"
	}

	return strings.Join([]string{"UserIdentifier", string(data)}, " ")
}
