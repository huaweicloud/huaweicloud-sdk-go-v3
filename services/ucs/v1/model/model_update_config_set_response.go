package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigSetResponse Response Object
type UpdateConfigSetResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateConfigSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigSetResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigSetResponse", string(data)}, " ")
}
