package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSecretVersionResponse struct {
	Version        *Version `json:"version,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowSecretVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretVersionResponse", string(data)}, " ")
}
