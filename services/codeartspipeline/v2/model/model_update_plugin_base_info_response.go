package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePluginBaseInfoResponse Response Object
type UpdatePluginBaseInfoResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePluginBaseInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginBaseInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdatePluginBaseInfoResponse", string(data)}, " ")
}
