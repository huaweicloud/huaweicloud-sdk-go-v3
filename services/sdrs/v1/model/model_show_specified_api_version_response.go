package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSpecifiedApiVersionResponse struct {
	Version        *ShowApiVersionParams `json:"version,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowSpecifiedApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpecifiedApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowSpecifiedApiVersionResponse", string(data)}, " ")
}
