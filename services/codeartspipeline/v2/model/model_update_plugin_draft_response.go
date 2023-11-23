package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePluginDraftResponse Response Object
type UpdatePluginDraftResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePluginDraftResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginDraftResponse struct{}"
	}

	return strings.Join([]string{"UpdatePluginDraftResponse", string(data)}, " ")
}
