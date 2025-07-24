package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIDcsResponse Response Object
type UpdateIDcsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIDcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIDcsResponse struct{}"
	}

	return strings.Join([]string{"UpdateIDcsResponse", string(data)}, " ")
}
