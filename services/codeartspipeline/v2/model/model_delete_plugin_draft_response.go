package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePluginDraftResponse Response Object
type DeletePluginDraftResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePluginDraftResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePluginDraftResponse struct{}"
	}

	return strings.Join([]string{"DeletePluginDraftResponse", string(data)}, " ")
}
