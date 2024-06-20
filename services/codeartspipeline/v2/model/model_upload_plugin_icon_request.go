package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadPluginIconRequest Request Object
type UploadPluginIconRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	PluginName string `json:"plugin_name"`

	Body *UploadPluginIconRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadPluginIconRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadPluginIconRequest struct{}"
	}

	return strings.Join([]string{"UploadPluginIconRequest", string(data)}, " ")
}
