package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchYamlTemplatesInfoRequest Request Object
type ShowAutoSearchYamlTemplatesInfoRequest struct {
}

func (o ShowAutoSearchYamlTemplatesInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchYamlTemplatesInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchYamlTemplatesInfoRequest", string(data)}, " ")
}
