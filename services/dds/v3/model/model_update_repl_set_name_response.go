package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReplSetNameResponse Response Object
type UpdateReplSetNameResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateReplSetNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplSetNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateReplSetNameResponse", string(data)}, " ")
}
