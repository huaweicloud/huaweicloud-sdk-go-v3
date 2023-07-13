package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowMetadataRespGesMetadataLabels struct {

	// Label名称。
	Name *string `json:"name,omitempty"`

	// 属性Map
	Properties *[]map[string]string `json:"properties,omitempty"`
}

func (o ShowMetadataRespGesMetadataLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadataRespGesMetadataLabels struct{}"
	}

	return strings.Join([]string{"ShowMetadataRespGesMetadataLabels", string(data)}, " ")
}
