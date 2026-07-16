package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchYamlTemplateContentResponse Response Object
type ShowAutoSearchYamlTemplateContentResponse struct {

	// yaml文件名称。
	FileName *string `json:"file_name,omitempty"`

	// yaml文件内容。
	Content        *string `json:"content,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAutoSearchYamlTemplateContentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchYamlTemplateContentResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchYamlTemplateContentResponse", string(data)}, " ")
}
