package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginSchema the request body of login
type LoginSchema struct {

	// ide_type
	IdeType string `json:"ide_type"`

	// ide_version
	IdeVersion string `json:"ide_version"`

	// plugin_version
	PluginVersion string `json:"plugin_version"`
}

func (o LoginSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginSchema struct{}"
	}

	return strings.Join([]string{"LoginSchema", string(data)}, " ")
}
