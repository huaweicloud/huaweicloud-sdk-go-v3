package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowReplSetNameResponse struct {
	Body *string `json:"body,omitempty"`

	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReplSetNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplSetNameResponse struct{}"
	}

	return strings.Join([]string{"ShowReplSetNameResponse", string(data)}, " ")
}
