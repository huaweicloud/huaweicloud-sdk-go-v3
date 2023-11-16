package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadBasicPluginResponse Response Object
type UploadBasicPluginResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadBasicPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadBasicPluginResponse struct{}"
	}

	return strings.Join([]string{"UploadBasicPluginResponse", string(data)}, " ")
}
