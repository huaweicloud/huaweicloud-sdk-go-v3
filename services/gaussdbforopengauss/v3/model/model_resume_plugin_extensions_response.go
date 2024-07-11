package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumePluginExtensionsResponse Response Object
type ResumePluginExtensionsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResumePluginExtensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePluginExtensionsResponse struct{}"
	}

	return strings.Join([]string{"ResumePluginExtensionsResponse", string(data)}, " ")
}
