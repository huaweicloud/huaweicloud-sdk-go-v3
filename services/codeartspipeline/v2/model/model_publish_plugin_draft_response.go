package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishPluginDraftResponse Response Object
type PublishPluginDraftResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishPluginDraftResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishPluginDraftResponse struct{}"
	}

	return strings.Join([]string{"PublishPluginDraftResponse", string(data)}, " ")
}
