package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadPluginIconResponse Response Object
type UploadPluginIconResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadPluginIconResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadPluginIconResponse struct{}"
	}

	return strings.Join([]string{"UploadPluginIconResponse", string(data)}, " ")
}
