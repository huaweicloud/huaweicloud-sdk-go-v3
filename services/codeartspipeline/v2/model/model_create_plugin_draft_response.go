package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginDraftResponse Response Object
type CreatePluginDraftResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePluginDraftResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginDraftResponse struct{}"
	}

	return strings.Join([]string{"CreatePluginDraftResponse", string(data)}, " ")
}
