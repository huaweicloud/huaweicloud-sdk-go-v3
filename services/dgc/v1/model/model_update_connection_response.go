package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateConnectionResponse", string(data)}, " ")
}
