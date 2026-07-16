package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchYamlTemplatesInfoResponse Response Object
type ShowAutoSearchYamlTemplatesInfoResponse struct {

	// 所有yaml文件的目录和文件名信息。
	YamlTemplates  *[]YamlTemplate `json:"yaml_templates,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowAutoSearchYamlTemplatesInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchYamlTemplatesInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchYamlTemplatesInfoResponse", string(data)}, " ")
}
