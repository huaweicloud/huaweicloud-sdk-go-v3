package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropertiesSchema struct {

	// context
	Context *string `json:"context,omitempty"`

	// docstring
	Docstring *string `json:"docstring,omitempty"`

	// the type of ide
	IdeType *string `json:"ide_type,omitempty"`

	// the version of ide
	IdeVersion *string `json:"ide_version,omitempty"`

	// code language
	Language string `json:"language"`

	// the version of plugin
	PluginVersion *string `json:"plugin_version,omitempty"`

	// signature
	Signature string `json:"signature"`
}

func (o PropertiesSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertiesSchema struct{}"
	}

	return strings.Join([]string{"PropertiesSchema", string(data)}, " ")
}
