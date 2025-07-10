package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstancePortResponse Response Object
type UpdateInstancePortResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateInstancePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstancePortResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstancePortResponse", string(data)}, " ")
}
