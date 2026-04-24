package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompleteResourceTokenAuthRequestBody struct {

	// Unique identifier for the user's authentication session (tracks OAuth2 flow state)
	SessionUri string `json:"session_uri"`

	UserIdentifier *UserIdentifier `json:"user_identifier"`
}

func (o CompleteResourceTokenAuthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompleteResourceTokenAuthRequestBody struct{}"
	}

	return strings.Join([]string{"CompleteResourceTokenAuthRequestBody", string(data)}, " ")
}
