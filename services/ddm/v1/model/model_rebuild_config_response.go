package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RebuildConfigResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebuildConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildConfigResponse struct{}"
	}

	return strings.Join([]string{"RebuildConfigResponse", string(data)}, " ")
}
