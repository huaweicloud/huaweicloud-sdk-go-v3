package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIRackResponse Response Object
type UpdateIRackResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIRackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIRackResponse struct{}"
	}

	return strings.Join([]string{"UpdateIRackResponse", string(data)}, " ")
}
